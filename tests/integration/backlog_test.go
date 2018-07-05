package integration

import (
	"os"
	"testing"

	"github.com/vvatanabe/go-backlog/backlog/shared"
	"github.com/vvatanabe/go-backlog/backlog/v2"
	"strconv"
)

var (
	clientV2 *v2.Client
	host     string
	apiKey   string
	projectID int
	issueTypeID int
	categoryID int
	versionID int
	assigneeID int
	parentIssueKey string
	parentIssueID int
)

func init() {
	host = os.Getenv("BACKLOG_HOST")
	apiKey = os.Getenv("BACKLOG_API_KEY")
	projectID, _ = strconv.Atoi(os.Getenv("BACKLOG_PROJECT_ID"))
	issueTypeID, _ = strconv.Atoi(os.Getenv("BACKLOG_ISSUE_TYPE_ID"))
	categoryID, _ = strconv.Atoi(os.Getenv("BACKLOG_CATEGORY_ID"))
	versionID, _ = strconv.Atoi(os.Getenv("BACKLOG_VERSION_ID"))
	assigneeID, _ = strconv.Atoi(os.Getenv("BACKLOG_ASSIGNEE_ID"))
	assigneeID, _ = strconv.Atoi(os.Getenv("BACKLOG_ASSIGNEE_ID"))
	parentIssueKey = os.Getenv("BACKLOG_PARENT_ISSUE_ID")
	parentIssueID, _ = strconv.Atoi(os.Getenv("BACKLOG_PARENT_ISSUE_ID"))

	if host == "" || apiKey == "" {
		panic("!!! Integration test using host and apiKey. !!!\n")
	} else {
		clientV2 = v2.NewClient(host, nil)
		clientV2.SetAPIKey(apiKey)
	}
}

func test(t *testing.T, result interface{}, resp *shared.Response, err error) {
	println("URL: " + resp.Request.URL.String())
	if err != nil {
		t.Fatalf("Returned error: %v", err)
	}
	if resp == nil {
		t.Fatal("Returned nil response")
	}
	if result == nil {
		t.Error("Returned nil result")
	}
}
