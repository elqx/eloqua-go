package bulk

import (
	"fmt"
	"context"
)

type CdosService service

type CdoField Field

type Cdo struct {
	Name string `json:name`
	DisplayNameFieldUri string `json:displayNameFieldUri`
	EmailAddressFieldUri string `json:emailAddressFieldUri`
	UniqueFieldUri string `json:uniqueFieldUri`
	Uri string `json:uri`
	CreatedBy string `json:createdBy`
	CreatedAt string `json:createdAt`
	UpdatedBy string `json:createdBy`
	UpdatedAt string `json:updatedAt`
}

type CdoFieldSearchResponse struct {
	Count int `json:"count,omitempty"`
	HasMore bool `json:"hasMore,omitempty"`
	Items []CdoField `json:"items,omitempty"`
	Limit int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	TotalResults int64 `json:"totalResults,omitempty"`
}

type CdoSearchResponse struct {
	Count int `json:"count,omitempty"`
	HasMore bool `json:"hasMore,omitempty"`
	Items []Cdo `json:"items,omitempty"`
	Limit int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	TotalResults int64 `json:"totalResults,omitempty"`
}

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-customobjects-parentid-exports-post.html
func (s *CdosService) CreateExport(ctx context.Context, parentId int, export *Export) (*Export, error) {
	u := fmt.Sprintf("/customObjects/%v/exports", parentId)
	req, err := s.client.NewRequest("POST", u, export)
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

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-customobjects-parentid-fields-get.html
func (s *CdosService) ListFields(ctx context.Context, parentId int) (*CdoFieldSearchResponse, error) {
	u := fmt.Sprintf("/customObjects/%v/fields", parentId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &CdoFieldSearchResponse{}
	_, err = s.client.Do(ctx, req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-customobjects-get.html
func (s *CdosService) List(ctx context.Context) (*CdoSearchResponse, error) {
	req, err := s.client.NewRequest("GET", "/customObjects", nil)
	if err != nil {
		return nil, err
	}

	r := &CdoSearchResponse{}
	_, err = s.client.Do(ctx, req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
