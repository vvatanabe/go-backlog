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

func Test_IssueService_AddIssue_should_add_a_issue(t *testing.T) {
	setup()
	defer teardown()

	projectID := 1
	summary := "summary"
	issueTypeID := 1
	priorityID := 2
	opt := &AddIssueOptions{
		ParentIssueID: 1,
		Description: "Description",
		StartDate: "2018-07-01",
		DueDate: "2018-07-02",
		EstimatedHours: 1.5,
		ActualHours: 2.0,
		CategoryIDs: []int{1},
		VersionIDs: []int{1},
		MilestoneIDs: []int{1},
		AssigneeID: 1,
		NotifiedUserIDs: []int{1},
		AttachmentIDs: []int{1},
		CustomFields: map[int]interface{}{
			1: "text",
			2: true,
			3: 1,
		},
		CustomFieldOtherValues: map[int]interface{}{
			1: "text",
			2: true,
			3: 1,
		},
	}

	b, _ := ioutil.ReadFile(fixturesPath + "add-issue.json")
	mux.HandleFunc("/issues",
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "POST")
			internal.TestFormValues(t, r, internal.Values{
				"projectId": projectID,
				"summary": summary,
				"issueTypeId": issueTypeID,
				"priorityId": priorityID,
				"parentIssueId": opt.ParentIssueID,
				"description": opt.Description,
				"startDate": opt.StartDate,
				"dueDate": opt.DueDate,
				"estimatedHours": opt.EstimatedHours,
				"actualHours": opt.ActualHours,
				"categoryId[]": opt.CategoryIDs[0],
				"versionId[]": opt.VersionIDs[0],
				"milestoneId[]": opt.MilestoneIDs[0],
				"assigneeId": opt.AssigneeID,
				"notifiedUserId[]": opt.NotifiedUserIDs[0],
				"attachmentId[]": opt.AttachmentIDs[0],
				"customField_1": opt.CustomFields[1],
				"customField_2": opt.CustomFields[2],
				"customField_3": opt.CustomFields[3],
				"customField_1_otherValue": opt.CustomFieldOtherValues[1],
				"customField_2_otherValue": opt.CustomFieldOtherValues[2],
				"customField_3_otherValue": opt.CustomFieldOtherValues[3],
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Issues.AddIssue(context.Background(), projectID, summary, issueTypeID, priorityID, opt)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *Issue
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}