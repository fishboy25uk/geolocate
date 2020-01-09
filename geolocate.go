package geolocate

import (
	"net"
	"strings"
)

//Geolocate performs a lookup for an IP address against the MaxMind GeoLite2 City and ASN databases and returns geolocation and ASN details in a Geolocation object
func Geolocate(ip net.IP) (Geolocation, error) {

	var recordCity struct {
		City struct {
			Names map[string]string `maxminddb:"names"`
		} `maxminddb:"city"`
		Country struct {
			ISOCode string            `maxminddb:"iso_code"`
			Names   map[string]string `maxminddb:"names"`
		} `maxminddb:"country"`
	} // Or any appropriate struct

	err := dbCity.Lookup(ip, &recordCity)
	if err != nil {
		return Geolocation{}, err
	}

	var recordASN struct {
		AutonomousSystemNumber       uint   `maxminddb:"autonomous_system_number"`
		AutonomousSystemOrganization string `maxminddb:"autonomous_system_organization"`
	}
	err = dbASN.Lookup(ip, &recordASN)
	if err != nil {
		return Geolocation{}, err
	}

	geolocation := Geolocation{City: recordCity.City.Names["en"], Country: recordCity.Country.Names["en"], CountryCode: strings.ToLower(recordCity.Country.ISOCode), ASN: recordASN.AutonomousSystemNumber, Org: recordASN.AutonomousSystemOrganization}

	return geolocation, nil

}
