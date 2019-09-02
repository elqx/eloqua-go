package bulk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestActivitiesService_CreateExport(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Export{
		Name: "export",
	}

	mux.HandleFunc("/activities/exports", func(w http.ResponseWriter, r *http.Request) {
		v := new(Export)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")

		want := &Export{Name: "export"}
		if !reflect.DeepEqual(v, want) {
			t.Errorf("Request body = %+v, want %+v", v, want)
		}

		fmt.Fprint(w, `{"uri":"/abc"}`)
	})

	ctx := context.Background()
	got, err := client.Activities.CreateExport(ctx, input)
	if err != nil {
		t.Errorf("Activities.CreateExport returned error: %v", err)
	}

	want := &Export{Uri: "/abc"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Activties.CreateExport returned %+v, want %+v", got, want)
	}
}

func TestActivitiesService_DeleteExport(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/activities/exports/1", func (w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	ctx := context.Background()
	_, err := client.Activities.DeleteExport(ctx, 1)
	if err != nil {
		t.Errorf("Activities.DeleteExport returned error: %v", err)
	}
}


func TestActivitiesService_DeleteExportData(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/activities/exports/1/data", func (w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	ctx := context.Background()
	_, err := client.Activities.DeleteExportData(ctx, 1)
	if err != nil {
		t.Errorf("Activities.DeleteExportDara returned error: %v", err)
	}
}

func TestActivitiesService_ListExports(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/activities/exports", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"items": [{"uri": "/abc"}]}`)
	})

	ctx := context.Background()
	got, err := client.Activities.ListExports(ctx)
	if err != nil {
		t.Errorf("Activities.ListExports returned error: %v", err)
	}

	want := &ActivityExportSearchResponse{Items: []Export{{Uri: "/abc"}}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Activties.ListExports returned %+v, want %+v", got, want)
	}
}

func TestActivitiesService_ListFields(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/activities/fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"items": [{"name": "abc"}]}`)
	})

	ctx := context.Background()
	got, err := client.Activities.ListFields(ctx, nil)
	if err != nil {
		t.Errorf("Activities.ListFields returned error: %v", err)
	}

	want := &ActivityFieldSearchResponse{Items: []ActivityField{{Name: "abc"}}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Activties.ListFields returned %+v, want %+v", got, want)
	}
}
