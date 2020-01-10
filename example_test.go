package geolocate

import (
	"fmt"
	"net"

	geolocate "github.com/fishboy25uk/geolocate"
)

func main() {

	//Update DBs
	geolocate.UpdateDBs("MAXMINDKEY")

	//Open DBs
	geolocate.OpenDBs()

	ip := net.ParseIP("9.9.9.9")

	g, _ := geolocate.Geolocate(ip)

	fmt.Printf("\nGeolocation Record for %v\nCity: %s\nCountry: %s\nCountry Code: %s\nASN: %v\nOrg: %s\n", ip, g.City, g.Country, g.CountryCode, g.ASN, g.Org)

}
