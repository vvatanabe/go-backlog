package v2

import (
	"net/http"
	"net/url"

	"github.com/vvatanabe/go-backlog/backlog/internal"
)

const (
	APIVersion = "v2"
)

type service struct {
	client *internal.Client
}

type Client struct {
	client *internal.Client

	Space        *SpaceService
	Projects     *ProjectsService
	Issues       *IssuesService
	PullRequests *PullRequestsService
}

func (c *Client) SetAPIKey(key string) {
	c.client.SetAPIKey(key)
}

func NewClient(host string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse("https://" + host + "/api/" + APIVersion + "/")

	c := &Client{client: internal.NewClient(baseURL, httpClient)}

	common := &service{client: c.client}

	c.Space = (*SpaceService)(common)
	c.Projects = (*ProjectsService)(common)
	c.Issues = (*IssuesService)(common)
	c.PullRequests = (*PullRequestsService)(common)

	return c
}
