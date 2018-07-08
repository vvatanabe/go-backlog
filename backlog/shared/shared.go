package shared

import (
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
)

type Response struct {
	*http.Response
}

type ErrorResponse struct {
	Response *http.Response
	Errors   []*Error `json:"errors"`
}

type Error struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"moreInfo"`
}

func (r *ErrorResponse) Error() string {
	bytes, _ := json.Marshal(r.Errors)
	return fmt.Sprintf("%v %v %d %v",
		r.Response.Request.Method, SanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, string(bytes))
}

func SanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("apiKey")) > 0 {
		params.Set("apiKey", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}
