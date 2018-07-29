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
	b, _ := ioutil.ReadFile(fixturesPath + "get-project-user-list.json")
	projectIDOrKey := "EXAMPLE"
	mux.HandleFunc(fmt.Sprintf("/projects/%s/users", projectIDOrKey),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Projects.GetProjectUsers(context.Background(), projectIDOrKey)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want []*User
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_ProjectsService_GetIssueTypes_should_get_some_issue_types_in_project(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-issue-type-list.json")
	projectIDOrKey := "EXAMPLE"
	mux.HandleFunc(fmt.Sprintf("/projects/%s/issueTypes", projectIDOrKey),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Projects.GetIssueTypes(context.Background(), projectIDOrKey)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want []*IssueType
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_ProjectsService_GetCategories_should_get_some_categories_in_project(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-issue-type-list.json")
	projectIDOrKey := "EXAMPLE"
	mux.HandleFunc(fmt.Sprintf("/projects/%s/categories", projectIDOrKey),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Projects.GetCategories(context.Background(), projectIDOrKey)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want []*Category
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_ProjectsService_GetVersions_should_get_some_versions_in_project(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-version-milestone-list.json")
	projectIDOrKey := "EXAMPLE"
	mux.HandleFunc(fmt.Sprintf("/projects/%s/versions", projectIDOrKey),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Projects.GetVersions(context.Background(), projectIDOrKey)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want []*Version
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_ProjectsService_GetCustomFields_should_get_some_custom_fields_in_project(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-custom-field-list.json")
	projectIDOrKey := "EXAMPLE"
	mux.HandleFunc(fmt.Sprintf("/projects/%s/customFields", projectIDOrKey),
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Projects.GetCustomFields(context.Background(), projectIDOrKey)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want []*CustomField
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_ProjectsService_GetPriorities_should_get_some_priority(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-priority-list.json")
	mux.HandleFunc("/priorities",
		func(w http.ResponseWriter, r *http.Request) {
			internal.TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Projects.GetPriorities(context.Background())
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want []*Priority
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}