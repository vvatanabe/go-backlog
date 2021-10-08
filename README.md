# go-backlog [![Build Status](https://travis-ci.org/vvatanabe/go-backlog.svg?branch=master)](https://travis-ci.org/vvatanabe/go-backlog) [![Coverage Status](https://coveralls.io/repos/github/vvatanabe/go-backlog/badge.svg?branch=master)](https://coveralls.io/github/vvatanabe/go-backlog?branch=master)

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

## Support API

- [x] [get space](https://developer.nulab-inc.com/ja/docs/backlog/api/2/get-space/)
- [x] [get space disk usage](https://developer.nulab.com/docs/backlog/api/2/get-space-disk-usage/#get-space-disk-usage)
- [x] [get issue](https://developer.nulab-inc.com/docs/backlog/api/2/get-issue/)
- [x] [add issue](https://developer.nulab-inc.com/docs/backlog/api/2/add-issue/)
- [x] [get project](https://developer.nulab-inc.com/docs/backlog/api/2/get-project/)
- [x] [get project user list](https://developer.nulab-inc.com/docs/backlog/api/2/get-project-user-list//)
- [x] [get issue type list](https://developer.nulab-inc.com/docs/backlog/api/2/get-issue-type-list/)
- [x] [add category list](https://developer.nulab-inc.com/docs/backlog/api/2/get-category-list/)
- [x] [get version milestone list](https://developer.nulab-inc.com/docs/backlog/api/2/get-version-milestone-list/)
- [x] [get custom field list](https://developer.nulab-inc.com/docs/backlog/api/2/get-custom-field-list/)
- [x] [get priority list](https://developer.nulab-inc.com/docs/backlog/api/2/get-priority-list/)
- [x] [get pull request](https://developer.nulab.com/ja/docs/backlog/api/2/get-pull-request/)
- [x] [add pull request comment](https://developer.nulab.com/docs/backlog/api/2/add-pull-request-comment/)
- [x] [get pull request comment](https://developer.nulab.com/docs/backlog/api/2/get-pull-request-comment/)