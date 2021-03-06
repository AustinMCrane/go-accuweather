package accuweather

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testKey = flag.String("test-key", "testtest", "accuweather test api key")

var wichitaKey = "348426"

func TestMain(m *testing.M) {
	flag.Parse()
	m.Run()
}

// read test json data from testdata directory
// used to load test responses from the iex api
func readTestData(fileName string) (string, error) {
	b, err := ioutil.ReadFile("test/responses/" + fileName)

	if err != nil {
		return "", err
	}

	str := string(b)
	return str, nil
}

type mockHTTPClient struct {
	body    string
	headers map[string]string
	code    int
	err     error
}

// Get is just to conform to HTTPClient
// it takes the mockHTTPClients and writes the responses for code, headers,
// body, headers ect...
func (c *mockHTTPClient) Get(url string) (*http.Response, error) {
	w := httptest.NewRecorder()
	w.WriteHeader(c.code)

	w.WriteString(c.body)

	for key, value := range c.headers {
		w.Header().Add(key, value)
	}

	resp := w.Result()
	return resp, c.err
}

func TestNewClient(t *testing.T) {
	c := NewClient(*testKey, &http.Client{})
	if c == nil {
		t.Fatalf("returned a nil client")
	}
}

func TestClientBadRequest(t *testing.T) {
	// 500 internal server error
	httpc := mockHTTPClient{
		body: "{}",
		code: http.StatusInternalServerError,
	}

	c := NewClient(*testKey, &httpc)
	_, err := c.CitySearch("wichita")

	if err == nil {
		t.Fatal("expected an internal server error")
	}
}

func TestCitySearch(t *testing.T) {
	body, err := readTestData("city_search.json")
	if err != nil {
		t.Fatal(err)
	}

	httpc := mockHTTPClient{body: body, code: 200}
	c := NewClient(*testKey, &httpc)
	result, err := c.CitySearch("wichita")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatalf("result was unexpectedly nil")
	}
}

func TestGeopositionSearch(t *testing.T) {
	body, err := readTestData("geoposition_search.json")
	if err != nil {
		t.Fatal(err)
	}

	httpc := mockHTTPClient{body: body, code: 200}
	c := NewClient(*testKey, &httpc)

	lat := 37.4256
	lon := -97.343460

	result, err := c.GeopositionSearch(lat, lon)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Fatalf("result was unexpectedly nil")
	}
}

func TestGetCurrentConditions(t *testing.T) {
	body, err := readTestData("current_conditions.json")
	if err != nil {
		t.Fatal(err)
	}

	httpc := mockHTTPClient{body: body, code: 200}
	c := NewClient(*testKey, &httpc)
	result, err := c.GetCurrentConditions(wichitaKey)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatalf("result was unexpectedly nil")
	}

	// no results
	httpc = mockHTTPClient{body: "[]", code: 200}
	c = NewClient(*testKey, &httpc)
	result, err = c.GetCurrentConditions(wichitaKey)
	if err != ErrNotFound {
		t.Fatalf("expected error not found")
	}
}

func TestGetDailyForecasts(t *testing.T) {
	ttable := []struct {
		Type        DailyForecastType
		LocationKey string
	}{
		{
			Type:        OneDay,
			LocationKey: wichitaKey,
		},
		{
			Type:        FiveDay,
			LocationKey: wichitaKey,
		},
		{
			Type:        TenDay,
			LocationKey: wichitaKey,
		},
		{
			Type:        FifteenDay,
			LocationKey: wichitaKey,
		},
	}

	body, err := readTestData("1day_forecast.json")
	if err != nil {
		t.Fatal(err)
	}

	for i, tt := range ttable {
		t.Run(fmt.Sprintf("test at index: %d", i), func(t *testing.T) {
			httpc := mockHTTPClient{body: body, code: 200}
			c := NewClient(*testKey, &httpc)
			result, err := c.GetDailyForecasts(tt.LocationKey, tt.Type)

			if err != nil {
				t.Fatal(err)
			}

			if result == nil {
				t.Fatalf("result was unexpectedly nil")
			}
		})
	}
}

func TestGetHourlyForecasts(t *testing.T) {
	ttable := []struct {
		Type        HourlyForecastType
		LocationKey string
	}{
		{
			Type:        OneHour,
			LocationKey: wichitaKey,
		},
		{
			Type:        TwelveHour,
			LocationKey: wichitaKey,
		},
		{
			Type:        TwentyFourHour,
			LocationKey: wichitaKey,
		},
		{
			Type:        SeventyTwoHour,
			LocationKey: wichitaKey,
		},
		{
			Type:        OneHundredTwentyHour,
			LocationKey: wichitaKey,
		},
	}

	body, err := readTestData("1hour_forecast.json")
	if err != nil {
		t.Fatal(err)
	}

	for i, tt := range ttable {
		t.Run(fmt.Sprintf("test at index: %d", i), func(t *testing.T) {
			httpc := mockHTTPClient{body: body, code: 200}
			c := NewClient(*testKey, &httpc)
			result, err := c.GetHourlyForecasts(tt.LocationKey, tt.Type)

			if err != nil {
				t.Fatal(err)
			}

			if result == nil {
				t.Fatalf("result was unexpectedly nil")
			}
		})
	}
}
