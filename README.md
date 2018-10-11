[![GoDoc](https://godoc.org/github.com/AustinMCrane/go-accuweather?status.svg)](https://godoc.org/github.com/AustinMCrane/go-accuweather)
[![Coverage Status](https://coveralls.io/repos/github/AustinMCrane/go-accuweather/badge.svg?branch=master)](https://coveralls.io/github/AustinMCrane/go-accuweather?branch=master)
[![Build Status](https://travis-ci.com/AustinMCrane/go-accuweather.svg?branch=master)](https://travis-ci.com/AustinMCrane/go-accuweather)
[![Go Report Card](https://goreportcard.com/badge/github.com/austinmcrane/go-accuweather)](https://goreportcard.com/badge/github.com/austinmcrane/go-accuweather)

### Current Features
- text based location lookup
- geo position location lookup via lat lon
- daily forecast
- hourly forecast
- current conditions (not historic)

#### Usage:

##### Current Conditions Example
```
package main

import (
  "fmt"
  "net/http"

  accuweather "github.com/AustinMCrane/go-accuweather"
)

func main() {
  client := accuweather.NewClient("APIKEY", &http.Client{})

  locations, err := client.CitySearch("LOCATION NAME")
  if err != nil {
    panic(err)
  }

  if len(locations) == 0 {
    fmt.Println(fmt.Sprintf("no locations found for %s", *location))
    return
  }

  locationKey := locations[0].Key
  conditions, err := client.GetCurrentConditions(locationKey)
  if err != nil {
    panic(err)
  }
}
```
