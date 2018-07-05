# go-backlog

go-backlog is a GO client library for accessing the [Backlog API](https://developer.nulab-inc.com/docs/backlog/).

## Installation

This package can be installed with the go get command:

```
$ go get github.com/vvatanabe/go-backlog
```

## Usage

### Import

``` go
import "github.com/vvatanabe/go-backlog/backlog/v2"
```

### Access APIs using API Key

```go
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
```

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)