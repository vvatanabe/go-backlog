package v2

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

const (
	fixturesPath = "../../testdata/v2/"
	// dummyAPIKey = "deadbeef"
	)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)


	parsedURL, _ := url.Parse(server.URL)
	client = NewClient("localhost",nil)
	client.client.BaseURL = parsedURL
	// client.SetAPIKey(dummyAPIKey)
}

func teardown() {
	server.Close()
}