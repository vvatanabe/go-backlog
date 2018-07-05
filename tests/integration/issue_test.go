package integration

import (
	"testing"
	"context"
	"github.com/vvatanabe/go-backlog/backlog/v2"
)

func Test_V2_Issue_GetIssue_should_get_a_issue(t *testing.T) {
	result, resp, err := clientV2.Issues.GetIssue(context.Background(), parentIssueKey)
	test(t, result, resp, err)
}

func Test_V2_Issue_AddIssue_should_add_a_issue(t *testing.T) {
	projectID := projectID
	summary := "integration test for go-backlog"
	issueTypeID := issueTypeID
	priorityID := 2
	assigneeID := assigneeID
	opt := &v2.AddIssueOptions{
		ParentIssueID: parentIssueID,
		Description: "hello",
		StartDate: "2018-07-01",
		DueDate: "2018-07-02",
		EstimatedHours: 1.5,
		ActualHours: 2.0,
		CategoryIDs: []int{categoryID},
		VersionIDs: []int{versionID},
		MilestoneIDs: []int{versionID},
		AssigneeID: assigneeID,
	}
	result, resp, err := clientV2.Issues.AddIssue(context.Background(), projectID, summary, issueTypeID, priorityID, opt)
	test(t, result, resp, err)
}