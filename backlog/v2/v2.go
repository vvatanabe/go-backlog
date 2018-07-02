package v2

import (
	"net/http"
	"net/url"

	"github.com/vvatanabe/go-backlog/backlog/internal"
)

const (
	ApiVersion = "v2"
)

type service struct {
	client *internal.Client
}

type Client struct {
	client *internal.Client

	Projects *ProjectsService
	Issues   *IssuesService
}

func (c *Client) SetAPIKey(key string) {
	c.client.SetAPIKey(key)
}

func NewClient(host string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse("https://" + host + "/api/" + ApiVersion + "/")

	c := &Client{client: internal.NewClient(baseURL, httpClient)}

	common := &service{client: c.client}

	c.Projects = (*ProjectsService)(common)
	c.Issues = (*IssuesService)(common)

	return c
}
