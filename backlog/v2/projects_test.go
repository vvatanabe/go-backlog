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

func Test_ProjectsService_GetProject_should_get_a_project(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-project.json")
	projectIDOrKey := "EXAMPLE"
	mux.HandleFunc(fmt.Sprintf("/projects/%s", projectIDOrKey),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Projects.GetProject(context.Background(), projectIDOrKey)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *Project
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_ProjectsService_GetProjectUsers_should_get_some_users_in_project(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-project.json")
	projectIDOrKey := "EXAMPLE"
	mux.HandleFunc(fmt.Sprintf("/projects/%s", projectIDOrKey),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Projects.GetProject(context.Background(), projectIDOrKey)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *Project
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

//func (s *ProjectsService) GetProjectUsers(ctx context.Context, projectIDOrKey string) ([]*User, *Response, error) {
//	u := fmt.Sprintf("projects/%s/users", projectIDOrKey)
//	var result []*User
//	if resp, err := s.client.Get(ctx, u, nil, &result); err != nil {
//		return nil, resp, err
//	} else {
//		return result, resp, nil
//	}
//}