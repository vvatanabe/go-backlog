package shared

import (
	"fmt"
	"net/http"
	"net/url"
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
	var msgs []string
	for _, err := range r.Errors {
		msg := fmt.Sprintf("message: %v, code: %v", err.Message, err.Code)
		msgs = append(msgs, msg)
	}
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, SanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, msgs)
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
