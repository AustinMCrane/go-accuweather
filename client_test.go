package accuweather

import (
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testKey = flag.String("test-key", "testtest", "accuweather test api key")

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
	_, err := c.SearchLocations("wichita")

	if err == nil {
		t.Fatal("expected an internal server error")
	}
}

func TestSearchLocations(t *testing.T) {
	body, err := readTestData("search_locations.json")
	if err != nil {
		t.Fatal(err)
	}

	httpc := mockHTTPClient{body: body, code: 200}
	c := NewClient(*testKey, &httpc)
	result, err := c.SearchLocations("wichita")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatalf("result was unexpectedly nil")
	}
}

func TestCurrentConditions(t *testing.T) {
	body, err := readTestData("current_conditions.json")
	if err != nil {
		t.Fatal(err)
	}

	httpc := mockHTTPClient{body: body, code: 200}
	c := NewClient(*testKey, &httpc)
	result, err := c.CurrentConditions("348426")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatalf("result was unexpectedly nil")
	}

	// no results
	httpc = mockHTTPClient{body: "[]", code: 200}
	c = NewClient(*testKey, &httpc)
	result, err = c.CurrentConditions("348426")
	if err != ErrNotFound {
		t.Fatalf("expected error not found")
	}

}
