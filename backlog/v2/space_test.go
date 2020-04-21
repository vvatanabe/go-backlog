package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/vvatanabe/go-backlog/backlog/internal"
)

func Test_SpaceService_GetSpace_should_get_a_space(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-space.json")
	mux.HandleFunc("/space",
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Space.GetSpace(context.Background())
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *Space
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_SpaceService_GetSpaceDiskUsage_should_get_a_disk_usage(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-space-disk-usage.json")
	mux.HandleFunc("/space/diskUsage",
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Space.GetSpaceDiskUsage(context.Background())
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *DiskUsage
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}
