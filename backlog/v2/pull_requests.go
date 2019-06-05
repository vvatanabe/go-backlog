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

// GetIssue docs: https://developer.nulab.com/ja/docs/backlog/api/2/get-pull-request/
func (s *PullRequestsService) GetPullRequest(ctx context.Context, projectIdOrKey, repoIdOrName string, number int) (*PullRequest, *shared.Response, error) {
	u := fmt.Sprintf("projects/%s/git/repositories/%s/pullRequests/%d", projectIdOrKey, repoIdOrName, number)
	var result *PullRequest
	resp, err := s.client.Get(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
