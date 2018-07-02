package httpc

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"sort"
)

type Response struct {
	*http.Response
}

func (resp *Response) DecodeJson(v interface{}) error {
	var reader io.ReadCloser
	var err error
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
	default:
		reader = resp.Body
	}
	defer reader.Close()

	decErr := json.NewDecoder(reader).Decode(v)
	if decErr == io.EOF {
		// ignore EOF errors caused by empty response body
		decErr = nil
	}
	if decErr != nil {
		return decErr
	}
	return nil
}

func (resp *Response) Copy(w io.Writer) (written int64, err error) {
	written, err = io.Copy(w, resp.Body)
	return
}

func (resp *Response) ReadAll() ([]byte, error) {
	var reader io.ReadCloser
	var err error
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
	default:
		reader = resp.Body
	}
	defer reader.Close()
	return ioutil.ReadAll(reader)
}

type Header map[string]string

func (h Header) Set(k, v string) {
	h[k] = v
}

func NewClient(client *http.Client) *Client {
	return &Client{
		client: client,
		Header: make(map[string]string),
	}
}

type Client struct {
	client       *http.Client
	Header       Header
	Query url.Values
}

func (c *Client) Post(ctx context.Context, url *url.URL, body url.Values) (*Response, error) {
	AddQuery(url, c.Query)
	return c.call(ctx, http.MethodPost, url, body)
}

func (c *Client) Put(ctx context.Context, url *url.URL, body url.Values) (*Response, error) {
	AddQuery(url, c.Query)
	return c.call(ctx, http.MethodPut, url, body)
}

func (c *Client) Delete(ctx context.Context, url *url.URL, query url.Values) (res *Response, err error) {
	AddQuery(url, c.Query, query)
	return c.call(ctx, http.MethodDelete, url, nil)
}

func (c *Client) Get(ctx context.Context, url *url.URL, query url.Values) (res *Response, err error) {
	AddQuery(url, c.Query, query)
	return c.call(ctx, http.MethodGet, url, nil)
}

func (c *Client) call(ctx context.Context, method string, url *url.URL, body url.Values) (*Response, error) {
	req, err := c.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *Client) newRequest(method string, url *url.URL, body url.Values) (*http.Request, error) {

	var buf io.Reader
	if body != nil {
		buf = strings.NewReader(body.Encode())
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	for k, v := range c.Header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

func (c *Client) newUploadRequest(url *url.URL, reader io.Reader, size int64, mediaType string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url.String(), reader)
	if err != nil {
		return nil, err
	}
	req.ContentLength = size

	for k, v := range c.Header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", mediaType)
	return req, nil

}

func (c *Client) do(ctx context.Context, req *http.Request) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	return &Response{Response: resp}, nil
}

func AddQuery(url *url.URL, queries... url.Values) {
	q := url.Values{}
	for _, query := range queries {
		keys := make([]string, 0, len(query))
		for k := range query {
			keys = append(keys, k)
		}
		for _, k := range keys {
			vs := query.Get(k)
			q.Add(k, vs)
		}
	}
	url.RawQuery = q.Encode()
}