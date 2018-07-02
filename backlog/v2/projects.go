package v2

import (
	"context"
	"fmt"

	. "github.com/vvatanabe/go-backlog/backlog/shared"
)

type ProjectsService service

type Project struct {
	ID                                int    `json:"id"`
	ProjectKey                        string `json:"projectKey"`
	Name                              string `json:"name"`
	ChartEnabled                      bool   `json:"chartEnabled"`
	SubtaskingEnabled                 bool   `json:"subtaskingEnabled"`
	ProjectLeaderCanEditProjectLeader bool   `json:"projectLeaderCanEditProjectLeader"`
	TextFormattingRule                string `json:"textFormattingRule"`
	Archived                          bool   `json:"archived"`
}

type User struct {
	ID           int           `json:"id"`
	UserID       string        `json:"userId"`
	Name         string        `json:"name"`
	RoleType     int           `json:"roleType"`
	Lang         string        `json:"lang"`
	MailAddress  string        `json:"mailAddress"`
	NulabAccount *NulabAccount `json:"nulabAccount"`
}

type NulabAccount struct {
	NulabID  string `json:"nulabId"`
	Name     string `json:"name"`
	UniqueID string `json:"uniqueId"`
}

type IssueType struct {
	ID           int    `json:"id"`
	ProjectID    int    `json:"projectId"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
}

type Category struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"displayOrder"`
}

type Milestone struct {
	ID             int    `json:"id"`
	ProjectID      int    `json:"projectId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	StartDate      string `json:"startDate"`
	ReleaseDueDate string `json:"releaseDueDate"`
	Archived       bool   `json:"archived"`
	DisplayOrder   int    `json:"displayOrder"`
}

type Version struct {
	ID             int    `json:"id"`
	ProjectID      int    `json:"projectId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	StartDate      string `json:"startDate"`
	ReleaseDueDate string `json:"releaseDueDate"`
	Archived       bool   `json:"archived"`
	DisplayOrder   int    `json:"displayOrder"`
}

// Backlog API docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-project/
func (s *ProjectsService) GetProject(ctx context.Context, projectIDOrKey string) (*Project, *Response, error) {
	u := fmt.Sprintf("projects/%s", projectIDOrKey)
	var result *Project
	if resp, err := s.client.Get(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// TODO apply excludeGroupMembers option
// Backlog API docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-project-user-list/
func (s *ProjectsService) GetProjectUsers(ctx context.Context, projectIDOrKey string) ([]*User, *Response, error) {
	u := fmt.Sprintf("projects/%s/users", projectIDOrKey)
	var result []*User
	if resp, err := s.client.Get(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// Backlog API docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-issue-type-list/
func (s *ProjectsService) GetIssueTypes(ctx context.Context, projectIdOrKey string) ([]*IssueType, *Response, error) {
	u := fmt.Sprintf("projects/%s/issueTypes", projectIdOrKey)
	var result []*IssueType
	if resp, err := s.client.Get(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// Backlog API docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-category-list/
func (s *ProjectsService) GetCategories(ctx context.Context, projectIdOrKey string) ([]*Category, *Response, error) {
	u := fmt.Sprintf("projects/%s/categories", projectIdOrKey)
	var result []*Category
	if resp, err := s.client.Get(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// Backlog API docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-version-milestone-list/
func (s *ProjectsService) GetVersions(ctx context.Context, projectIdOrKey string) ([]*Version, *Response, error) {
	u := fmt.Sprintf("projects/%s/versions", projectIdOrKey)
	var result []*Version
	if resp, err := s.client.Get(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}
