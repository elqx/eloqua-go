package bulk

import (
	"context"
	"fmt"
	"net/http"
)

type AccountsService service

type AccountField Field

type AccountFieldSearchResponse struct {
	Count        int            `json:"count,omitempty"`
	HasMore      bool           `json:"hasMore,omitempty"`
	Items        []AccountField `json:"items,omitempty"`
	Limit        int            `json:"limit,omitempty"`
	Offset       int            `json:"offset,omitempty"`
	TotalResults int64          `json:"totalResults,omitempty"`
}

func (s *AccountsService) CreateExport(ctx context.Context, export *Export) (*Export, error) {
	req, err := s.client.NewRequest("POST", "/accounts/exports", export)
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

func (s *AccountsService) DeleteExport(ctx context.Context, id int) (*http.Response, error) {
	u := fmt.Sprintf("/accounts/exports/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

func (s *AccountsService) DeleteExportData(ctx context.Context, id int) (*http.Response, error) {
	u := fmt.Sprintf("/accounts/exports/%v/data", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *AccountsService) GetFields(ctx context.Context) (*AccountFieldSearchResponse, error) {
	req, err := s.client.NewRequest("GET", "/accounts/fields", nil)
	if err != nil {
		return nil, err
	}

	r := &AccountFieldSearchResponse{}
	_, err = s.client.Do(ctx, req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
