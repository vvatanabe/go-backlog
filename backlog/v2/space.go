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
