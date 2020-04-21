package v2

import (
	"context"
	"time"

	"github.com/vvatanabe/go-backlog/backlog/shared"
)

type SpaceService service

type Space struct {
	SpaceKey           string    `json:"spaceKey"`
	Name               string    `json:"name"`
	OwnerID            int       `json:"ownerId"`
	Lang               string    `json:"lang"`
	Timezone           string    `json:"timezone"`
	ReportSendTime     string    `json:"reportSendTime"`
	TextFormattingRule string    `json:"textFormattingRule"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated"`
}

// GetSpace docs: https://developer.nulab-inc.com/ja/docs/backlog/api/2/get-space/
func (s *SpaceService) GetSpace(ctx context.Context) (*Space, *shared.Response, error) {
	u := "space"
	var result *Space
	resp, err := s.client.Get(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type DiskUsage struct {
	Capacity   int                `json:"capacity"`
	Issue      int                `json:"issue"`
	Wiki       int                `json:"wiki"`
	File       int                `json:"file"`
	Subversion int                `json:"subversion"`
	Git        int                `json:"git"`
	GitLFS     int                `json:"gitLFS"`
	Details    []*DiskUsageDetail `json:"details"`
}

type DiskUsageDetail struct {
	ProjectID  int `json:"projectId"`
	Issue      int `json:"issue"`
	Wiki       int `json:"wiki"`
	File       int `json:"file"`
	Subversion int `json:"subversion"`
	Git        int `json:"git"`
	GitLFS     int `json:"gitLFS"`
}

// GetSpace docs: https://developer.nulab.com/docs/backlog/api/2/get-space-disk-usage/
func (s *SpaceService) GetSpaceDiskUsage(ctx context.Context) (*DiskUsage, *shared.Response, error) {
	u := "/space/diskUsage"
	var result *DiskUsage
	resp, err := s.client.Get(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
