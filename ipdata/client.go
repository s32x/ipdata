package ipdata /* import "s32x.com/ipdata/ipdata" */

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"

	maxminddb "github.com/oschwald/maxminddb-golang"
	"golang.org/x/sync/errgroup"
)

// Client contains the databases needed to get info on a particular IP address
type Client struct {
	mu        sync.Mutex
	city, asn maxminddb.Reader
}

// NewClient creates and returns a fully populated ipdata Client
func NewClient() (*Client, error) {
	c := &Client{mu: sync.Mutex{}}
	if err := c.updateReaders(); err != nil {
		return nil, err
	}
	return c, nil
}

// updateReaders concurrently retrieves and populates the databases using the
// urls on the Client
func (c *Client) updateReaders() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Concurrently retrieve and populate the mmdb databases
	var g errgroup.Group
	g.Go(readMMDB(geoLite2City, &c.city))
	g.Go(readMMDB(geoLite2ASN, &c.asn))
	return g.Wait()
}

// Close performs the final task of closing the maxminddb Readers
func (c *Client) Close() {
	c.city.Close()
	c.asn.Close()
}

// readMMDB takes a compressed MaxMind database URL, downloads, decompresses,
// and generates a maxmindb.Reader
func readMMDB(url string, out *maxminddb.Reader) func() error {
	return func() error {
		// Retrieve the compressed database from the url
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		// GZIP read the response body
		gzr, err := gzip.NewReader(res.Body)
		if err != nil {
			return err
		}
		defer gzr.Close()

		// Create a reader for reading all tarred files
		tr := tar.NewReader(gzr)
		for {
			// Untar the next file
			header, err := tr.Next()
			if err != nil {
				if err == io.EOF {
					return errors.New("No files found with mmdb extention")
				}
				return err
			}

			// If it's a file and it has the extension mmdb
			if header.Typeflag == tar.TypeReg &&
				strings.Contains(header.Name, ".mmdb") {
				// Create a new bytes buffer and load the reader into it
				buf := new(bytes.Buffer)
				buf.ReadFrom(tr)

				// Create a db from the buffered Bytes and set it on the out
				db, err := maxminddb.FromBytes(buf.Bytes())
				if err != nil {
					return err
				}
				*out = *db
				return nil
			}
		}
	}
}
