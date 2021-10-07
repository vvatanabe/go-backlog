package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/vvatanabe/go-backlog/backlog/internal"
)

func Test_PullRequestsService_GetPullRequest_should_get_a_pull_request(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-pull-request.json")
	projectIdOrKey := "FOO"
	issueIDOrKey := "EXAMPLE-1"
	number := 1
	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s/pullRequests/%d", projectIdOrKey, issueIDOrKey, number),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.PullRequests.GetPullRequest(context.Background(), projectIdOrKey, issueIDOrKey, number)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *PullRequest
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_PullRequestsService_AddPullRequest_should_get_a_pull_request(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "add-pull-request.json")
	projectIdOrKey := "FOO"
	issueIDOrKey := "EXAMPLE-1"

	opt := &AddPullRequestOptions{
		Summary:     "foo",
		Description: "bar",
		Base:        "master",
		Branch:      "topic",
	}
	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s/pullRequests", projectIdOrKey, issueIDOrKey),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "POST")
			internal.TestFormValues(t, r, internal.Values{
				"summary":     opt.Summary,
				"description": opt.Description,
				"base":        opt.Base,
				"branch":      opt.Branch,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.PullRequests.AddPullRequest(context.Background(), projectIdOrKey, issueIDOrKey, opt)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *PullRequest
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_PullRequestsService_AddPullRequestComment_should_get_a_pull_request_comment(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "add-pull-request-comment.json")
	projectIdOrKey := "FOO"
	repoIDOrKey := "EXAMPLE-1"
	number := 1

	opt := &AddPullRequestCommentOptions{
		Content:        "hello",
		NotifiedUserID: []int{1},
		AttachmentID:   []int{2},
	}
	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s/pullRequests/%d/comments", projectIdOrKey, repoIDOrKey, number),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "POST")
			internal.TestFormValues(t, r, internal.Values{
				"content":          opt.Content,
				"attachmentId[]":   opt.AttachmentID[0],
				"notifiedUserId[]": opt.NotifiedUserID[0],
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.PullRequests.AddPullRequestComment(context.Background(), projectIdOrKey, repoIDOrKey, number, opt)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *PullRequestComment
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}
