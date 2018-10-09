package accuweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

const baseEndpoint = "http://dataservice.accuweather.com"

// ErrNotFound 404 error
var ErrNotFound = errors.New("api returned no results")

// HTTPClient an interface to describe simple requests to a url
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

// Client accuweather http request client
type Client struct {
	apiKey   string
	client   HTTPClient
	Language string
}

// NewClient create a new accuweather http request client with api key
func NewClient(apiKey string, httpClient HTTPClient) *Client {
	return &Client{
		apiKey:   apiKey,
		client:   httpClient,
		Language: "en-us", // change after creating a new client
	}
}

// AccuAPIRequest is a base object for any accuweather api request
// includes the api key to every request
type AccuAPIRequest struct {
	APIKey   string `url:"apikey"`   // api key from accuweather console
	Language string `url:"language"` // what language the data will be returned in ie: 'en-us'
}

func (c *Client) newAccuRequest() *AccuAPIRequest {
	return &AccuAPIRequest{
		APIKey:   c.apiKey,
		Language: c.Language,
	}
}

// CitySearch returns a list of cities found with a search query
// example: `CitySearch("new york")` will return new your new york
// as one of the results
//
// accuweather api docs:
// https://developer.accuweather.com/accuweather-locations-api/apis/get/locations/v1/cities/search
func (c *Client) CitySearch(search string) ([]*Location, error) {
	accuRequest := c.newAccuRequest()
	req := &searchRequest{
		AccuAPIRequest: *accuRequest,
		Query:          search,
	}

	var result []*Location
	err := c.getJSON("/locations/v1/cities/search", req, &result)
	return result, err
}

// GeopositionSearch returns a city closest to the latitude and longitude pair
//
// accuweather api docs:
// https://developer.accuweather.com/accuweather-locations-api/apis/get/locations/v1/cities/geoposition/search
func (c *Client) GeopositionSearch(lat float64, lon float64) (*Location, error) {

	// lat lon is a comma seperated list lat,lon
	latlon := fmt.Sprintf("%f,%f", lat, lon)

	accuRequest := c.newAccuRequest()
	req := &searchRequest{
		AccuAPIRequest: *accuRequest,
		Query:          latlon,
	}

	var result *Location
	err := c.getJSON("/locations/v1/cities/geoposition/search", req, &result)
	return result, err
}

type searchRequest struct {
	AccuAPIRequest
	Query string `url:"q,omitempty"`
}

// GetCurrentConditions gets the current conditions for a location by location key,
// get location keys from SearchLocation or other location search functions.
//
// accuweather api docs:
// https://developer.accuweather.com/accuweather-current-conditions-api/apis/get/currentconditions/v1/%7BlocationKey%7D
func (c *Client) GetCurrentConditions(locationKey string) (*CurrentCondition, error) {
	req := c.newAccuRequest()
	var result []*CurrentCondition
	err := c.getJSON("/currentconditions/v1/"+locationKey, req, &result)
	if len(result) == 0 {
		return nil, ErrNotFound
	}

	// NOTE: not sure why this api returns an array, i think it shouldnt
	return result[0], err
}

// GetDailyForecasts gets 1day, 5day, 10day, or 15day forecast for the location key
//
// accuweather api docs:
// https://developer.accuweather.com/accuweather-forecast-api/apis/get/forecasts/v1/daily/1day/%7BlocationKey%7D
func (c *Client) GetDailyForecasts(locationKey string, forecastType DailyForecastType) (*Forecast, error) {
	req := c.newAccuRequest()
	var result Forecast
	err := c.getJSON("/forecasts/v1/daily/"+forecastType.String()+"/"+locationKey, req, &result)
	return &result, err
}

// getJSON a generic way to send a get request and marshal response into a interface
func (c *Client) getJSON(route string, request interface{}, response interface{}) error {
	url := c.endpoint(route)

	values, err := query.Values(request)
	if err != nil {
		return err
	}

	queryString := values.Encode()
	if queryString != "" {
		url = url + "?" + queryString
	}

	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%v: %v", resp.Status, string(body))
	}

	dec := json.NewDecoder(resp.Body)
	return dec.Decode(response)
}

// endpoint returns the concatinated version of the url with the base
// url in front of the route
func (c *Client) endpoint(route string) string {
	return baseEndpoint + route
}
