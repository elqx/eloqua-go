package bulk


import (
	"context"
	"fmt"
	"net/http"
)

type ContactsService service

func (s *ContactsService) CreateExport(ctx context.Context, export *Export) (*Export, error) {
	req, err := s.client.NewRequest("POST", "/contacts/exports", export)
	if err != nil {
		return nil, err
	}

	r := &Export{}
	_, err = s.client.Do(ctx, req, r)
	if err != nil {
		return nil, err
	}

	return r, nil

}

func (s *ContactsService) DeleteExport(ctx context.Context, id int) (*http.Response, error) {
	u := fmt.Sprintf("/contacts/exports/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

func (s *ContactsService) DeleteExportData(ctx context.Context, id int) (*http.Response, error) {
	u := fmt.Sprintf("/contacts/exports/%v/data", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
