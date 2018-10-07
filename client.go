package accuweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

const baseEndpoint = "http://dataservice.accuweather.com"

// HTTPClient an interface to describe simple requests to a url
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

// Client accuweather http request client
type Client struct {
	apiKey string
	client HTTPClient
}

// NewClient create a new accuweather http request client with api key
func NewClient(apiKey string, httpClient HTTPClient) *Client {
	return &Client{
		apiKey: apiKey,
		client: httpClient,
	}
}

// AccuAPIRequest is a base object for any accuweather api request
// includes the api key to every request
type AccuAPIRequest struct {
	APIKey string `url:"apiKey"`
}

func (c *Client) NewAccuRequest() *AccuAPIRequest {
	return &AccuAPIRequest{
		APIKey: c.apiKey,
	}
}

func (c *Client) SearchForLocation(search string) ([]*Location, error) {
	accuRequest := c.NewAccuRequest()
	req := &searchLocationsRequest{
		AccuAPIRequest: *accuRequest,
		Query:          search,
	}

	log.Println(req)
	var result []*Location
	err := c.getJSON("/locations/v1/cities/search", req, &result)
	return result, err
}

type searchLocationsRequest struct {
	AccuAPIRequest
	Query string `url:"q,omitempty"`
}

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

func (c *Client) endpoint(route string) string {
	return baseEndpoint + route
}
