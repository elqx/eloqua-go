package bulk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCdosService_CreateExport(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Export{
		Name: "export",
	}

	mux.HandleFunc("/customObjects/1/exports", func(w http.ResponseWriter, r *http.Request) {
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
	got, err := client.Cdos.CreateExport(ctx, 1, input)
	if err != nil {
		t.Errorf("Cdos.CreateExport returned error: %v", err)
	}

	want := &Export{Uri: "/abc"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Cdos.CreateExport returned %+v, want %+v", got, want)
	}
}


func TestCdosService_ListFields(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customObjects/1/fields", func (w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"items": [{"name": "Email Address"}]}`)
	})

	ctx := context.Background()
	got, err := client.Cdos.ListFields(ctx, 1)
	if err != nil {
		t.Errorf("Cdos.ListFields returned error: %v", err)
	}

	want := &CdoFieldSearchResponse{Items: []CdoField{{Name: "Email Address"}}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Cdo.ListFields returned %+v, want %+v", got, want)
	}

}
