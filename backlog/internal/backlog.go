package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	. "github.com/vvatanabe/go-backlog/backlog/shared"
	"github.com/vvatanabe/go-backlog/httpc"
)

const (
	LibraryVersion   = "0.9.0"
	UserAgent        = "go-backlog/" + LibraryVersion
	DefaultMediaType = "application/octet-stream"
)

func NewClient(baseURL *url.URL, httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := httpc.NewClient(httpClient)
	c.Header.Set("User-Agent", UserAgent)
	c.Header.Set("Content-Type", DefaultMediaType)

	return &Client{
		client:  c,
		baseURL: baseURL,
	}
}

type Client struct {
	client  *httpc.Client
	baseURL *url.URL
	apiKey  string
}

func (c *Client) SetAPIKey(key string) {
	c.client.Query.Set("apiKey", key)
}

type RequestFunc func(context.Context, *url.URL, url.Values) (*httpc.Response, error)

func (c *Client) do(ctx context.Context, uri string, p, v interface{}, request RequestFunc) (*Response, error) {

	rel, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	values, err := toValues(p)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	resp, err := request(ctx, u, values)
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			if parsedURL, err := url.Parse(e.URL); err == nil {
				e.URL = SanitizeURL(parsedURL).String()
				return nil, e
			}
		}
		return nil, err
	}
	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

	err = resp.DecodeJson(v)
	if err != nil {
		return nil, err
	}

	return &Response{Response: resp.Response}, nil
}

func (c *Client) Post(ctx context.Context, uri string, body interface{}, v interface{}) (*Response, error) {
	return c.do(ctx, uri, body, v, c.client.Post)
}

func (c *Client) Put(ctx context.Context, uri string, body interface{}, v interface{}) (*Response, error) {
	return c.do(ctx, uri, body, v, c.client.Put)
}

func (c *Client) Delete(ctx context.Context, uri string, q, v interface{}) (*Response, error) {
	return c.do(ctx, uri, q, v, c.client.Delete)
}

func (c *Client) Get(ctx context.Context, uri string, q, v interface{}) (*Response, error) {
	return c.do(ctx, uri, q, v, c.client.Get)
}

func checkResponse(r *httpc.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{}
	if err := r.DecodeJson(errorResponse); err != nil {
		return err
	} else {
		errorResponse.Response = r.Response
		return errorResponse
	}
}

func toValues(data interface{}) (url.Values, error) {
	result := make(map[string]interface{})
	b, _ := json.Marshal(data)
	d := json.NewDecoder(strings.NewReader(string(b)))
	d.UseNumber()
	if err := d.Decode(&result); err != nil {
		return nil, err
	}
	values := url.Values{}
	for k, v := range result {
		addValues(values, k, v)
	}
	return values, nil
}

func addValues(values url.Values, k string, v interface{}) {
	if as, ok := v.([]interface{}); ok {
		for i, v := range as {
			addValues(values, fmt.Sprintf(k, i), v)
		}
	} else {
		values.Add(k, fmt.Sprintf("%v", v))
	}
}
