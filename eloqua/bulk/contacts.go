package bulk


import (
	"context"
	"fmt"
	"net/http"
)

type ContactsService service


type ContactField struct {
	Name string `json:"name"`
	InternalName string `json:"internalName"`
	DataType string `json:"dataType"`
	HasReadOnlyConstraint bool `json:"hasReadOnlyConstrainti,omitempty"`
	HasNotNullConstraint bool `json:"hasNotNullConstraint,omitempty"`
	HasUniquenessConstraint bool `json:"hasUniquenessConstraint,omitempty"`
	Statement string `json:"statement"`
	Uri string `json:"uri"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ContactFieldSearchResponse struct {
	Count int `json:"count,omitempty"`
	HasMore bool `json:"hasMore,omitempty"`
	Items []ContactField `json:"items,omitempty"`
	Limit int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	TotalResults int64 `json:"totalResults,omitempty"`
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
