package v2

import (
	"time"
	"context"

	. "github.com/vvatanabe/go-backlog/backlog/shared"
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

// Backlog API docs: https://developer.nulab-inc.com/ja/docs/backlog/api/2/get-space/
func (s *SpaceService) GetSpace(ctx context.Context) (*Space, *Response, error) {
	u := "space"
	var result *Space
	if resp, err := s.client.Get(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}