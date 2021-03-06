# geolocate
Helper functions for IP geolocation in Go using the MaxMind Geolite2 databases.

Automatically downloads/updates MaxMind Geolite2 databases (City and ASN) on local machine (when UpdateDBs function is run) then loads DBs into memory for performance. Download requires provision of a MaxMind license key (free for GeoLite2 DBs but requires registration)

Geolocate function returns a GeoIP object with City, Country, CountryCode, ASN and Org name.

Uses [https://www.github.com/oschwald/maxminddb-golang](https://www.github.com/oschwald/maxminddb-golang)

NB: City / Country names only returned in English.

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
	geolocate.UpdateDBs("MAXMINDKEY")

	//Open DBs
	geolocate.OpenDBs()

	//Define IP for lookup
	ip := net.ParseIP("9.9.9.9")

	//Perform geolocation on IP (
	g, _ := geolocate.Geolocate(ip)

	fmt.Printf("\nGeolocation Record for %v\nCity: %s\nCountry: %s\nCountry Code: %s\nASN: %v\nOrg: %s\n", ip, g.City, g.Country, g.CountryCode, g.ASN, g.Org)
	
	//Geolocation Record for 9.9.9.9
	//City: Paris
	//Country: France
	//Country Code: fr
	//ASN: 19281
	//Org: Quad9

}

```
