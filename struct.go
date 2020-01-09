package geolocate

//Geolocation defines the struct for an IP geolocation record
type Geolocation struct {
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countrycode,omitempty"`
	ASN         uint   `json:"asn,omitempty"`
	Org         string `json:"org,omitempty"`
}
