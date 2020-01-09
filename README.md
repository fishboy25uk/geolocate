# geolocate
Helper functions for IP geolocation using MaxMind Geolite2 databases in Go

Automatically downloads/updates MaxMind Geolite2 database on local machine (when UpdateDBs is run) then loads DBs into memory for performance. Download required provision of a MaxMind license key (free for GeoLite2 DBs but required registration)

Lookup returns a GeoIP object with City, Country, CountryCode, ASN and Org name

Uses [https://www.github.com/oschwald/maxminddb-golang](https://www.github.com/oschwald/maxminddb-golang)

## Installation
```
go get -u github.com/fishboy25uk/geolocate
```
##Usage
```
package main

import github.com/fishboy25uk/geolocate

func main() {

  //Update DBs
	geolocate.UpdateDBs("MAXMIND KEY")

  //Open DBs
  geolocate.OpenDBs()

  g,_:=geolocate.Geolocate(net.ParseIP("8.8.8.8")

}

```
