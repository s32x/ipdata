package ipdata

import (
	"net"
	"strings"

	"github.com/mmcloughlin/geohash"
)

// IPData is a struct that contains information about a particular IP address
type IPData struct {
	IPAddress   string  `json:"ip_address,omitempty"`
	Hostname    string  `json:"hostname,omitempty"`
	ISP         string  `json:"isp,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	CountryName string  `json:"country_name,omitempty"`
	RegionCode  string  `json:"region_code,omitempty"`
	RegionName  string  `json:"region_name,omitempty"`
	City        string  `json:"city,omitempty"`
	ZipCode     string  `json:"zip_code,omitempty"`
	TimeZone    string  `json:"time_zone,omitempty"`
	GeoHash     string  `json:"geohash,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	MetroCode   uint    `json:"metro_code,omitempty"`
}

// Lookup performs the task of retrieving and returning the users IP address
// info
func (c *Client) Lookup(ipStr string) *IPData {
	// Reverse lookup the passed IP to retrieve any hostname that exists
	var hostname string
	if hs, _ := net.LookupAddr(ipStr); len(hs) > 0 {
		hostname = strings.TrimSuffix(hs[0], ".")
	}

	// Parse the IP address string passed
	ip := net.ParseIP(ipStr)

	// Lock and defer unlock the client to prevent concurrent read/write
	c.mu.Lock()
	defer c.mu.Unlock()

	// Lookup the Geo and ASN info from the stored client dbs
	var asn ASN
	var city City
	c.asn.Lookup(ip, &asn)
	c.city.Lookup(ip, &city)

	// Populate and return a fully populated ipdata
	t := IPData{
		IPAddress:   ip.String(),
		Hostname:    hostname,
		ISP:         asn.Organization,
		CountryCode: city.Country.ISOCode,
		CountryName: city.Country.Names["en"],
		City:        city.City.Names["en"],
		ZipCode:     city.Postal.Code,
		TimeZone:    city.Location.TimeZone,
		Latitude:    city.Location.Latitude,
		Longitude:   city.Location.Longitude,
		MetroCode:   city.Location.MetroCode,
		GeoHash: geohash.Encode(city.Location.Latitude,
			city.Location.Longitude),
	}
	if len(city.Region) > 0 {
		t.RegionCode = city.Region[0].ISOCode
		t.RegionName = city.Region[0].Names["en"]
	}
	return &t
}
