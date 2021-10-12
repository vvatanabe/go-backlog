package v2

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/vvatanabe/go-backlog/backlog/internal"
)

const (
	APIVersion     = "v2"
	ApiVersionPath = "api/" + APIVersion + "/"
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

func buildBaseURL(urlStr string) (*url.URL, error) {
	urlStr = strings.Replace(urlStr, "http://", "https://", 1)
	if !strings.HasPrefix(urlStr, "https://") {
		urlStr = "https://" + urlStr
	}

	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(baseURL.Path, ApiVersionPath) {
		baseURL.Path += ApiVersionPath
	}

	return baseURL, nil
}

func NewClient(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, _ := buildBaseURL(baseURL)

	c := &Client{client: internal.NewClient(u, httpClient)}

	common := &service{client: c.client}

	c.Space = (*SpaceService)(common)
	c.Projects = (*ProjectsService)(common)
	c.Issues = (*IssuesService)(common)
	c.PullRequests = (*PullRequestsService)(common)

	return c
}
