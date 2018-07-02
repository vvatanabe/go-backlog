package v2

import (
	"testing"
	"net/http"
	"fmt"
	"context"
	"encoding/json"
	"reflect"
	"io/ioutil"
)

func Test_SpaceService_GetSpace_should_get_a_space(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-space.json")
	mux.HandleFunc("/api/v2/space",
		func(w http.ResponseWriter, r *http.Request) {
			//TestMethod(t, r, "GET")
			//TestQueryValues(t, r, Values{"excludesGuest": true})
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
