package httpc

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"os"
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

func (resp *Response) String() (string, error) {
	v, err := resp.ReadAll()
	if err != nil {
		return "", err
	}
	return string(v), nil
}

type Header map[string]string

func (h Header) Set(k, v string) {
	h[k] = v
}

func NewClient(client *http.Client) *Client {
	return &Client{
		client:    client,
		Header:    make(map[string]string),
		BaseQuery: url.Values{},
	}
}

type Client struct {
	client    *http.Client
	Header    Header
	BaseQuery url.Values
}

func (c *Client) Post(ctx context.Context, u *url.URL, body url.Values) (*Response, error) {
	return c.call(ctx, http.MethodPost, c.resolveURL(u), body)
}

func (c *Client) Put(ctx context.Context, u *url.URL, body url.Values) (*Response, error) {
	return c.call(ctx, http.MethodPut, c.resolveURL(u), body)
}

func (c *Client) Delete(ctx context.Context, u *url.URL, query url.Values) (res *Response, err error) {
	return c.call(ctx, http.MethodDelete, c.resolveURL(u, query), nil)
}

func (c *Client) Get(ctx context.Context, u *url.URL, query url.Values) (res *Response, err error) {
	return c.call(ctx, http.MethodGet, c.resolveURL(u, query), nil)
}

func (c *Client) call(ctx context.Context, method, url string, body url.Values) (*Response, error) {
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

func (c *Client) newRequest(method, url string, body url.Values) (*http.Request, error) {

	var buf io.Reader
	if body != nil {
		buf = strings.NewReader(body.Encode())
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	for k, v := range c.Header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

func (c *Client) newUploadRequest(url string, reader io.Reader, size int64, mediaType string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, reader)
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

func (c *Client) NewMultipartRequest(url string, values map[string]io.Reader) (*http.Request, error) {
	var buffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&buffer)
	for key, reader := range values {
		var fieldWriter io.Writer
		var err error = nil
		if closable, ok := reader.(io.Closer); ok {
			closable.Close()
		}
		if file, ok := reader.(*os.File); ok {
			if fieldWriter, err = multipartWriter.CreateFormFile(key, file.Name()); err != nil {
				return nil, err
			}
		} else {
			if fieldWriter, err = multipartWriter.CreateFormField(key); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(fieldWriter, reader); err != nil {
			return nil, err
		}
	}
	multipartWriter.Close()
	req, err := http.NewRequest(http.MethodPost, url, &buffer)
	if err != nil {
		return nil, err
	}
	for k, v := range c.Header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
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

func (c *Client) resolveURL(u *url.URL, queries ...url.Values) string {
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
	u.RawQuery = q.Encode()
	baseQuery := c.BaseQuery.Encode()

	// TODO
	if baseQuery != "" {
		if u.RawQuery == "" {
			return u.String() + "?" + baseQuery
		} else {
			return u.String() + "&" + baseQuery
		}
	}

	return u.String()
}
