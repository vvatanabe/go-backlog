package v2

import (
	"context"
	"fmt"

	"github.com/vvatanabe/go-backlog/backlog/shared"
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

// GetProject docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-project/
func (s *ProjectsService) GetProject(ctx context.Context, projectIDOrKey string) (*Project, *shared.Response, error) {
	u := fmt.Sprintf("projects/%s", projectIDOrKey)
	var result *Project
	resp, err := s.client.Get(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetProjectUsers docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-project-user-list/
// TODO apply excludeGroupMembers option
func (s *ProjectsService) GetProjectUsers(ctx context.Context, projectIDOrKey string) ([]*User, *shared.Response, error) {
	u := fmt.Sprintf("projects/%s/users", projectIDOrKey)
	var result []*User
	resp, err := s.client.Get(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetIssueTypes docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-issue-type-list/
func (s *ProjectsService) GetIssueTypes(ctx context.Context, projectIDOrKey string) ([]*IssueType, *shared.Response, error) {
	u := fmt.Sprintf("projects/%s/issueTypes", projectIDOrKey)
	var result []*IssueType
	resp, err := s.client.Get(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetCategories docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-category-list/
func (s *ProjectsService) GetCategories(ctx context.Context, projectIDOrKey string) ([]*Category, *shared.Response, error) {
	u := fmt.Sprintf("projects/%s/categories", projectIDOrKey)
	var result []*Category
	resp, err := s.client.Get(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetVersions docs: https://developer.nulab-inc.com/docs/backlog/api/2/get-version-milestone-list/
func (s *ProjectsService) GetVersions(ctx context.Context, projectIDOrKey string) ([]*Version, *shared.Response, error) {
	u := fmt.Sprintf("projects/%s/versions", projectIDOrKey)
	var result []*Version
	resp, err := s.client.Get(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
