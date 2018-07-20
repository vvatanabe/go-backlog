package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/vvatanabe/go-backlog/backlog/shared"
	"github.com/vvatanabe/go-backlog/httpc"
)

const (
	Version          = "0.9.1"
	UserAgent        = "go-backlog/" + Version
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
		BaseURL: baseURL,
	}
}

type Client struct {
	client  *httpc.Client
	BaseURL *url.URL
}

func (c *Client) SetAPIKey(key string) {
	c.client.BaseQuery.Set("apiKey", key)
}

type RequestFunc func(context.Context, *url.URL, url.Values) (*httpc.Response, error)

func (c *Client) do(ctx context.Context, uri string, p, v interface{}, request RequestFunc) (*shared.Response, error) {

	rel, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	values, err := toValues(p)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	resp, err := request(ctx, u, values)
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			if parsedURL, err := url.Parse(e.URL); err == nil {
				e.URL = shared.SanitizeURL(parsedURL).String()
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

	return &shared.Response{Response: resp.Response}, nil
}

func (c *Client) Post(ctx context.Context, uri string, body interface{}, v interface{}) (*shared.Response, error) {
	return c.do(ctx, uri, body, v, c.client.Post)
}

func (c *Client) Put(ctx context.Context, uri string, body interface{}, v interface{}) (*shared.Response, error) {
	return c.do(ctx, uri, body, v, c.client.Put)
}

func (c *Client) Delete(ctx context.Context, uri string, q, v interface{}) (*shared.Response, error) {
	return c.do(ctx, uri, q, v, c.client.Delete)
}

func (c *Client) Get(ctx context.Context, uri string, q, v interface{}) (*shared.Response, error) {
	return c.do(ctx, uri, q, v, c.client.Get)
}

func checkResponse(r *httpc.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &shared.ErrorResponse{}
	if err := r.DecodeJson(errorResponse); err != nil {
		return err
	}
	errorResponse.Response = r.Response
	return errorResponse
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
	switch as := v.(type) {
	case []interface{}:
		for _, v := range as {
			addValues(values, k, v)
		}
	case map[string]interface{}:
		for mapKey, v := range as {
			addValues(values, fmt.Sprintf(k, mapKey), v)
		}
	default:
		values.Add(k, fmt.Sprintf("%v", v))
	}
}
