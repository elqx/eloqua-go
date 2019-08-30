package bulk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestContactsService_CreateExport(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Export{
		Name: "export",
	}

	mux.HandleFunc("/contacts/exports", func(w http.ResponseWriter, r *http.Request) {
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
	got, err := client.Contacts.CreateExport(ctx, input)
	if err != nil {
		t.Errorf("Contacts.CreateExport returned error: %v", err)
	}

	want := &Export{Uri: "/abc"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Contacts.CreateExport returned %+v, want %+v", got, want)
	}
}
/*
func TestContactsService_DeleteExport(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactss/exports/1", func (w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	ctx := context.Background()
	_, err := client.Activities.DeleteExport(ctx, 1)
	if err != nil {
		t.Errorf("Activities.DeleteExport returned error: %v", err)
	}
}


func TestContactsService_DeleteExportData(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contacts/exports/1/data", func (w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	ctx := context.Background()
	_, err := client.Activities.DeleteExportData(ctx, 1)
	if err != nil {
		t.Errorf("Activities.DeleteExportDara returned error: %v", err)
	}
}
*/
