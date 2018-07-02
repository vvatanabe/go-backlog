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

func Test_IssueService_GetIssue_should_get_a_issue(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-issue.json")
	issueIdOrKey := "EXAMPLE-1"
	mux.HandleFunc(fmt.Sprintf("/issues/%s", issueIdOrKey),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Issues.GetIssue(context.Background(), issueIdOrKey)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *Issue
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}
