package main

import (
	"context"
	"fmt"

	"github.com/vvatanabe/go-backlog/backlog/v2"
)

func main() {
	client := v2.NewClient("example.backlog.jp", nil)
	client.SetAPIKey("yourAPIKey")
	ctx := context.Background()
	issueIDOrKey := "EXAMPLE-1"
	issue, resp, err := client.Issues.GetIssue(ctx, issueIDOrKey)
	if err != nil {
		fmt.Printf("code: %v, err: %v", resp.StatusCode, err.Error())
	}
	fmt.Printf("value: %v", issue)
}
