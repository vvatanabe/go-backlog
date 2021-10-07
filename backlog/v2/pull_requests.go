package v2

import (
	"context"
	"fmt"
	"time"

	"github.com/vvatanabe/go-backlog/backlog/shared"
)

type PullRequestsService service

type PullRequest struct {
	ID           int           `json:"id"`
	ProjectID    int           `json:"projectId"`
	RepositoryID int           `json:"repositoryId"`
	Number       int           `json:"number"`
	Summary      string        `json:"summary"`
	Description  string        `json:"description"`
	Base         string        `json:"base"`
	Branch       string        `json:"branch"`
	Status       Status        `json:"status"`
	Assignee     User          `json:"assignee"`
	Issue        Issue         `json:"issue"`
	BaseCommit   string        `json:"baseCommit"`
	BranchCommit string        `json:"branchCommit"`
	CloseAt      time.Time     `json:"closeAt"`
	MergeAt      time.Time     `json:"mergeAt"`
	CreatedUser  User          `json:"createdUser"`
	Created      time.Time     `json:"created"`
	UpdatedUser  User          `json:"updatedUser"`
	Updated      time.Time     `json:"updated"`
	Attachments  []*Attachment `json:"attachments"`
	Stars        []*Star       `json:"stars"`
}

// GetPullRequest docs: https://developer.nulab.com/ja/docs/backlog/api/2/get-pull-request/
func (s *PullRequestsService) GetPullRequest(ctx context.Context, projectIdOrKey, repoIdOrName string, number int) (*PullRequest, *shared.Response, error) {
	u := fmt.Sprintf("projects/%s/git/repositories/%s/pullRequests/%d", projectIdOrKey, repoIdOrName, number)
	var result *PullRequest
	resp, err := s.client.Get(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type AddPullRequestOptions struct {
	Summary        string `json:"summary,omitempty"`
	Description    string `json:"description,omitempty"`
	Base           string `json:"base,omitempty"`
	Branch         string `json:"branch,omitempty"`
	IssueID        int    `json:"issueId,omitempty"`
	AssigneeID     int    `json:"assigneeId,omitempty"`
	NotifiedUserID []int  `json:"notifiedUserId,omitempty"`
	AttachmentID   []int  `json:"attachmentId,omitempty"`
}

// AddPullRequest docs: https://developer.nulab.com/ja/docs/backlog/api/2/add-pull-request/
func (s *PullRequestsService) AddPullRequest(ctx context.Context, projectIdOrKey, repoIdOrName string, opt *AddPullRequestOptions) (*PullRequest, *shared.Response, error) {
	u := fmt.Sprintf("projects/%s/git/repositories/%s/pullRequests", projectIdOrKey, repoIdOrName)
	var result *PullRequest
	resp, err := s.client.Post(ctx, u, opt, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type AddPullRequestCommentOptions struct {
	Content        string `json:"content"`
	NotifiedUserID []int  `json:"notifiedUserId[],omitempty"`
	AttachmentID   []int  `json:"attachmentId[],omitempty"`
}

type PullRequestComment struct {
	ID            int           `json:"id"`
	Content       string        `json:"content"`
	ChangeLog     []*ChangeLog  `json:"changeLog"`
	CreatedUser   *User         `json:"createdUser"`
	Created       *time.Time    `json:"created"`
	Updated       *time.Time    `json:"updated"`
	Stars         []interface{} `json:"stars"`
	Notifications []interface{} `json:"notifications"`
}

type ChangeLog struct {
	Field         string      `json:"field"`
	NewValue      string      `json:"newValue"`
	OriginalValue interface{} `json:"originalValue"`
}

// AddPullRequestComment docs: https://developer.nulab.com/docs/backlog/api/2/add-pull-request-comment/
func (s *PullRequestsService) AddPullRequestComment(ctx context.Context, projectIdOrKey, repoIdOrName string, number int, opt *AddPullRequestCommentOptions) (*PullRequestComment, *shared.Response, error) {
	u := fmt.Sprintf("projects/%s/git/repositories/%s/pullRequests/%d/comments", projectIdOrKey, repoIdOrName, number)
	var result *PullRequestComment
	resp, err := s.client.Post(ctx, u, opt, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
