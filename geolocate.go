package geolocate

import (
	"net"
	"strings"
)

//Geolocate performs a lookup for an IP address against the MaxMind GeoLite2 City and ASN databases and returns geolocation and ASN details in a Geolocation object
func Geolocate(ip net.IP) (Geolocation, error) {

	//ip := net.ParseIP(ipRaw)
	recordCity, err := dbCity.City(ip)
	if err != nil {
		return Geolocation{}, err
	}
	recordASN, err := dbASN.ASN(ip)
	if err != nil {
		return Geolocation{}, err
	}

	geoip := Geolocation{City: recordCity.City.Names["en"], Country: recordCity.Country.Names["en"], CountryCode: strings.ToLower(recordCity.Country.IsoCode), ASN: recordASN.AutonomousSystemNumber, ASNOrg: recordASN.AutonomousSystemOrganization}

	return geoip, nil

}
