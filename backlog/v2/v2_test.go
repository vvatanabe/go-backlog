package v2

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

const (
	fixturesPath = "../../testdata/v2/"

// dummyAPIKey = "deadbeef"
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	parsedURL, _ := url.Parse(server.URL)
	client = NewClient("localhost", nil)
	client.client.SetBaseURL(parsedURL)
	// client.SetAPIKey(dummyAPIKey)
}

func teardown() {
	server.Close()
}

func Test_buildBaseURL(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			name: "pattern 1",
			args: args{
				urlStr: "foo.backlog.com",
			},
			want: func() *url.URL {
				v, _ := url.Parse("https://foo.backlog.com/api/v2/")
				return v
			}(),
			wantErr: false,
		},
		{
			name: "pattern 2",
			args: args{
				urlStr: "foo.backlog.com/",
			},
			want: func() *url.URL {
				v, _ := url.Parse("https://foo.backlog.com/api/v2/")
				return v
			}(),
			wantErr: false,
		},
		{
			name: "pattern 3",
			args: args{
				urlStr: "https://foo.backlog.com",
			},
			want: func() *url.URL {
				v, _ := url.Parse("https://foo.backlog.com/api/v2/")
				return v
			}(),
			wantErr: false,
		},
		{
			name: "pattern 4",
			args: args{
				urlStr: "https://foo.backlog.com/",
			},
			want: func() *url.URL {
				v, _ := url.Parse("https://foo.backlog.com/api/v2/")
				return v
			}(),
			wantErr: false,
		},
		{
			name: "pattern 5",
			args: args{
				urlStr: "https://foo.backlog.com/api/v2",
			},
			want: func() *url.URL {
				v, _ := url.Parse("https://foo.backlog.com/api/v2/")
				return v
			}(),
			wantErr: false,
		},
		{
			name: "pattern 6",
			args: args{
				urlStr: "https://foo.backlog.com/api/v2/",
			},
			want: func() *url.URL {
				v, _ := url.Parse("https://foo.backlog.com/api/v2/")
				return v
			}(),
			wantErr: false,
		},
		{
			name: "pattern 7",
			args: args{
				urlStr: "http://foo.backlog.com",
			},
			want: func() *url.URL {
				v, _ := url.Parse("https://foo.backlog.com/api/v2/")
				return v
			}(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildBaseURL(tt.args.urlStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildBaseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildBaseURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
