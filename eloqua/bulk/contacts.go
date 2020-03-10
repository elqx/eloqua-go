package bulk

import (
	"context"
	"fmt"
	"net/http"
)

type ContactsService service

type ContactField Field

type LeadModel struct {
	Name      string  `json:"name,omitempty"`
	Status    string  `json:"status,omitempty"`
	Id        int     `json:"id,omitempty"`
	Fields    []Field `json:"fields,omitempty"`
	Uri       string  `json":uri,omitempty"`
	CreatedAt string  `json:"createdAt,omitempty"`
	CreatedBy string  `json:"createdBy,omitempty"`
	UpdatedAt string  `json:"updatedAt,omitempty"`
	UpdatedBy string  `json:"updatedBy,omitempty"`
}

type ContactFieldSearchResponse struct {
	Count        int            `json:"count,omitempty"`
	HasMore      bool           `json:"hasMore,omitempty"`
	Items        []ContactField `json:"items,omitempty"`
	Limit        int            `json:"limit,omitempty"`
	Offset       int            `json:"offset,omitempty"`
	TotalResults int64          `json:"totalResults,omitempty"`
}

type LeadModelsSearchResponse struct {
	Count        int         `json:"count,omitempty"`
	HasMore      bool        `json:"hasMore,omitempty"`
	Items        []LeadModel `json:"items,omitempty"`
	Limit        int         `json:"limit,omitempty"`
	Offset       int         `json:"offset,omitempty"`
	TotalResults int64       `json:"totalResults,omitempty"`
}

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

func (s *ContactsService) GetFields(ctx context.Context) (*ContactFieldSearchResponse, error) {
	req, err := s.client.NewRequest("GET", "/contacts/fields", nil)
	if err != nil {
		return nil, err
	}

	r := &ContactFieldSearchResponse{}
	_, err = s.client.Do(ctx, req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *ContactsService) ListLeadModels(ctx context.Context) (*LeadModelsSearchResponse, error) {
	req, err := s.client.NewRequest("GET", "/contacts/scoring/models", nil)
	if err != nil {
		return nil, err
	}

	r := &LeadModelsSearchResponse{}
	_, err = s.client.Do(ctx, req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
