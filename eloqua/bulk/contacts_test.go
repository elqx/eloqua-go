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

func TestContactsService_DeleteExport(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contacts/exports/1", func (w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	ctx := context.Background()
	_, err := client.Contacts.DeleteExport(ctx, 1)
	if err != nil {
		t.Errorf("Contacts.DeleteExport returned error: %v", err)
	}
}


func TestContactsService_DeleteExportData(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contacts/exports/1/data", func (w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	ctx := context.Background()
	_, err := client.Contacts.DeleteExportData(ctx, 1)
	if err != nil {
		t.Errorf("Contacts.DeleteExportDara returned error: %v", err)
	}
}

func TestContactsService_GetFields(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contacts/fields", func (w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"items": [{"name": "Email Address"}]}`)
	})

	ctx := context.Background()
	got, err := client.Contacts.GetFields(ctx)
	if err != nil {
		t.Errorf("Contacts.GetFields returned error: %v", err)
	}

	want := &ContactFieldSearchResponse{Items: []ContactField{{Name: "Email Address"}}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Contacts.GetFields returned %+v, want %+v", got, want)
	}

}

func TestContactsService_ListLeadModels(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contacts/scoring/models", func (w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"items": [{"name": "LeadModel1"}]}`)
	})

	ctx := context.Background()
	got, err := client.Contacts.ListLeadModels(ctx)
	if err != nil {
		t.Errorf("Contacts.ListLeadModels returned error: %v", err)
	}

	want := &LeadModelsSearchResponse{Items: []LeadModel{{Name: "LeadModel1"}}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Contacts.ListLeadModels returned %+v, want %+v", got, want)
	}

}
