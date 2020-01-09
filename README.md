# geolocate
Helper functions for IP geolocation using MaxMind Geolite2 databases in Go

Uses [https://www.github.com/oschwald/geoip2-golang](https://www.github.com/oschwald/geoip2-golang)

## Installation
```
go get -u github.com/fishboy25uk/geolocate
```
##Usage
```
package main

import github.com/fishboy25uk/geolocate

func main() {

  g,_:=geolocate.Geolocate(net.ParseIP("8.8.8.8")

}

```
