package v2

import (
	"context"
	"fmt"
	"time"

	. "github.com/vvatanabe/go-backlog/backlog/shared"
)

type IssuesService service

type Priority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Assignee struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	RoleType    int         `json:"roleType"`
	Lang        interface{} `json:"lang"`
	MailAddress string      `json:"mailAddress"`
}

type Attachment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Size int    `json:"size"`
}

type Star struct {
	ID        int         `json:"id"`
	Comment   interface{} `json:"comment"`
	URL       string      `json:"url"`
	Title     string      `json:"title"`
	Presenter *User       `json:"presenter"`
	Created   *time.Time  `json:"created"`
}

type Resolution struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Issue struct {
	ID             int           `json:"id"`
	ProjectID      int           `json:"projectId"`
	IssueKey       string        `json:"issueKey"`
	KeyID          int           `json:"keyId"`
	IssueType      *IssueType    `json:"issueType"`
	Summary        string        `json:"summary"`
	Description    string        `json:"description"`
	Resolutions    *Resolution   `json:"resolutions"`
	Priority       *Priority     `json:"priority"`
	Status         *Status       `json:"status"`
	Assignee       *Assignee     `json:"assignee"`
	Category       []*Category   `json:"category"`
	Versions       []interface{} `json:"versions"`
	Milestone      []*Milestone  `json:"milestone"`
	StartDate      string        `json:"startDate"`
	DueDate        string        `json:"dueDate"`
	EstimatedHours float64       `json:"estimatedHours"`
	ActualHours    float64       `json:"actualHours"`
	ParentIssueID  *int          `json:"parentIssueId"`
	CreatedUser    *User         `json:"createdUser"`
	Created        *time.Time    `json:"created"`
	UpdatedUser    *User         `json:"updatedUser"`
	Updated        *time.Time    `json:"updated"`
	CustomFields   []interface{} `json:"customFields"`
	Attachments    []*Attachment `json:"attachments"`
	SharedFiles    []interface{} `json:"sharedFiles"`
	Stars          []*Star       `json:"stars"`
}

type AddIssueOptions struct {
	ParentIssueID   int     `json:"parentIssueId,omitempty"`
	Description     string  `json:"description,omitempty"`
	StartDate       string  `json:"startDate,omitempty"`
	DueDate         string  `json:"dueDate,omitempty"`
	EstimatedHours  float64 `json:"estimatedHours,omitempty"`
	ActualHours     float64 `json:"actualHours,omitempty"`
	CategoryIDs     []int   `json:"categoryId[],omitempty"`
	VersionIDs      []int   `json:"versionId[],omitempty"`
	MilestoneIDs    []int   `json:"milestoneId[],omitempty"`
	AssigneeID      int     `json:"assigneeId,omitempty"`
	NotifiedUserIDs []int   `json:"notifiedUserId[],omitempty"`
	AttachmentIDs   []int   `json:"attachmentId[],omitempty"`
}

type addIssueOptions struct {
	ProjectID   int    `json:"projectId,omitempty"`
	Summary     string `json:"summary,omitempty"`
	IssueTypeID int    `json:"issueTypeId,omitempty"`
	PriorityID  int    `json:"priorityId,omitempty"`
	*AddIssueOptions
}

// Backlog API docs: https://developer.nulab-inc.com/docs/backlog/api/2/add-issue/
func (s *IssuesService) AddIssue(ctx context.Context, projectID int, summary string, issueTypeID, priorityID int, opt *AddIssueOptions) (*Issue, *Response, error) {
	u := "issues"
	body := &addIssueOptions{
		ProjectID:       projectID,
		Summary:         summary,
		IssueTypeID:     issueTypeID,
		PriorityID:      priorityID,
		AddIssueOptions: opt,
	}
	var result *Issue
	if resp, err := s.client.Post(ctx, u, body, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// Backlog API docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-issue/
func (s *IssuesService) GetIssue(ctx context.Context, issueIDOrKey string) (*Issue, *Response, error) {
	u := fmt.Sprintf("issues/%s", issueIDOrKey)
	var result *Issue
	if resp, err := s.client.Get(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}
