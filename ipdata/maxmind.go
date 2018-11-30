package ipdata

const (
	maxMindDB       = "http://geolite.maxmind.com/download/geoip/database/"
	geoLite2City    = maxMindDB + "GeoLite2-City.tar.gz"
	geoLite2Country = maxMindDB + "GeoLite2-Country.tar.gz"
	geoLite2ASN     = maxMindDB + "GeoLite2-ASN.tar.gz"
)

// ASN is the query response for ASN lookups
type ASN struct {
	Number       uint64 `maxminddb:"autonomous_system_number"`
	Organization string `maxminddb:"autonomous_system_organization"`
}

// City is the default query used for database lookups.
type City struct {
	Continent struct {
		Names map[string]string `maxminddb:"names"`
	} `maxminddb:"continent"`
	Country struct {
		ISOCode string            `maxminddb:"iso_code"`
		Names   map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	Region []struct {
		ISOCode string            `maxminddb:"iso_code"`
		Names   map[string]string `maxminddb:"names"`
	} `maxminddb:"subdivisions"`
	City struct {
		Names map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Location struct {
		Latitude  float64 `maxminddb:"latitude"`
		Longitude float64 `maxminddb:"longitude"`
		MetroCode uint    `maxminddb:"metro_code"`
		TimeZone  string  `maxminddb:"time_zone"`
	} `maxminddb:"location"`
	Postal struct {
		Code string `maxminddb:"code"`
	} `maxminddb:"postal"`
}
