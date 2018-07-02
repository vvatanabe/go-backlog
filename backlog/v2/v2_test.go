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

const fixturesPath = "../../testdata/v2/"

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	url, err := url.Parse(server.URL)
	if err != nil {
		panic(err.Error())
	}
	
	client = NewClient(url.Host,nil)
}

func teardown() {
	server.Close()
}