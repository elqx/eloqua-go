package bulk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSyncsService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Sync{
		SyncedInstanceURI: "/abc",
	}

	mux.HandleFunc("/syncs", func(w http.ResponseWriter, r *http.Request) {
		v := new(Sync)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")

		want := &Sync{SyncedInstanceURI: "/abc"}
		if !reflect.DeepEqual(v, want) {
			t.Errorf("Request body = %+v, want %+v", v, want)
		}

		fmt.Fprint(w, `{"uri":"/syncs/1"}`)
	})

	ctx := context.Background()
	got, err := client.Syncs.Create(ctx, input)
	if err != nil {
		t.Errorf("Syncs.Create returned error: %v", err)
	}

	want := &Sync{Uri: "/syncs/1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Syncs.Create returned %+v, want %+v", got, want)
	}
}

func TestSyncsService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/syncs/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"uri":"/syncs/1"}`)
	})

	ctx := context.Background()
	got, err := client.Syncs.Get(ctx, 1)
	if err != nil {
		t.Errorf("Syncs.Get returned error: %v", err)
	}

	want := &Sync{Uri: "/syncs/1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Syncs.Get returned %+v, want %+v", got, want)
	}
}

func TestSyncsService_GetData(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/syncs/1/data", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"items": [{"id": "123"}]}`)
	})

	ctx := context.Background()
	opt := &QueryOptions{Offset: 1000}
	got, err := client.Syncs.GetData(ctx, 1, opt)
	if err != nil {
		t.Errorf("Syncs.GetData returned error: %v", err)
	}

	item := Item{"id": "123"}
	items := []Item{item}
	want := &SyncDataQueryResponse{Items: items}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Syncs.GetData returned %+v, want %+v", got, want)
	}
}
