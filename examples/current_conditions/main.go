package main

import (
	"flag"
	"fmt"
	"net/http"

	accuweather "github.com/AustinMCrane/go-accuweather"
)

var location = flag.String("l", "wichita", "location to get current conditions for")
var key = flag.String("k", "", "accuweather api key")

func printConditions(location *accuweather.Location,
	conditions *accuweather.CurrentCondition) {
	if conditions == nil {
		return
	}

	fmt.Println(location.LocalizedName)
	fmt.Println(fmt.Sprintf("Temperature: %.2ff", conditions.Temperature.Imperial.Value))
	fmt.Println(conditions.WeatherText)
}

func main() {
	flag.Parse()
	client := accuweather.NewClient(*key, &http.Client{})

	locations, err := client.CitySearch(*location)
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

	printConditions(locations[0], conditions)
}
