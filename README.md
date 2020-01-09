# geolocate
Helper functions for IP geolocation using MaxMind Geolite2 databases in Go

Automatically downloads/updates MaxMind Geolite2 database on local machine (when UpdateDBs is run) then loads DBs into memory for performance. Download required provision of a MaxMind license key (free for GeoLite2 DBs but required registration)

Lookup returns a GeoIP object with City, Country, CountryCode, ASN and Org name

Uses [https://www.github.com/oschwald/maxminddb-golang](https://www.github.com/oschwald/maxminddb-golang)

## Installation
```
go get -u github.com/fishboy25uk/geolocate
```
## Usage
```
package main

import (
	"fmt"
	"net"

	"github.com/fishboy25uk/geolocate"
)

func main() {

	//Update DBs
	geolocate.UpdateDBs("jWYLmVxv2BMfplzT")

	//Open DBs
	geolocate.OpenDBs()

	ip := "9.9.9.9"

	g, _ := geolocate.Geolocate(net.ParseIP(ip))

	fmt.Printf("\nGeolocation Record for %s\nCity: %s\nCountry: %s\nCountry Code: %s\nASN: %v\nOrg: %s\n", ip, g.City, g.Country, g.CountryCode, g.ASN, g.Org)
	
	//Geolocation Record for 9.9.9.9
	//City: Paris
	//Country: France
	//Country Code: fr
	//ASN: 19281
	//Org: Quad9

}

```
