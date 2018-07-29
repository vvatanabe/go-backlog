package integration

import (
	"testing"
	"context"
)

func Test_V2_Space_GetPriorities_should_get_a_space(t *testing.T) {
	result, resp, err := clientV2.Projects.GetPriorities(context.Background())
	test(t, result, resp, err)
}