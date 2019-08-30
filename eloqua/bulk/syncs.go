package bulk

import (
	"context"
	"fmt"
)

type SyncsService service

type Sync struct {
	CallbackURL string `json:"callbackUrl,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	Status string `json:"status,omitempty"`
	SyncedInstanceURI string `json:"syncedInstanceUri,omitempty"`
	SyncedEndAt string `json:"syncedEndAt,omitempty"`
	SyncStartedAt string `json:"syncedStartedAt,omitempty"`
	Uri string `json:"uri,omitempty"`
}

type SyncLog struct {
	SyncURI string `json:"syncUri"`
	Count int `json:"count"`
	Severity string `json:"serverity"`
	StatusCode string `json:"statusCode"`
	Message string `json:"message"`
	CreatedAt string `json:"createdAt"`
}

type Item map[string]string

type SyncDataQueryResponse struct {
	Count int `json:"count,omitempty"`
	HasMore bool `json:"hasMore,omitempty"`
	Items []Item `json:"items,omitempty"`
	Limit int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	TotalResults int64 `json:"totalResults,omitempty"`
}

type QueryOptions struct {
	// Specifies the maximum number of records to return.
	// If not specified, the default it 1000.
	Limit int `url:"limit,omitempty"`

	Links string `url:"links,omitempty"`

	// Specifies an offset that allows you to retrieve
	// the next batch of records. Any positive integer.
	// If not specified the default is 0.
	Offset int `url:"offset,omitempty"`

	// Specified whether Total number of sync results found should be returned.
	// If not specified, the default is true.
	TotalResults bool `url:"totalResults,omitempty"`
}

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-syncs-post.html
func (s *SyncsService) Create(ctx context.Context, sync *Sync) (*Sync, error) {
	req, err := s.client.NewRequest("POST", "/syncs", sync)
	if err != nil {
		return nil, err
	}

	r := &Sync{}
	_, err = s.client.Do(ctx, req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-syncs-id-get.html
func (s *SyncsService) Get(ctx context.Context, id int) (*Sync, error) {
	u := fmt.Sprintf("/syncs/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &Sync{}
	_, err = s.client.Do(ctx, req, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-syncs-id-data-get.html
func (s *SyncsService) GetData(ctx context.Context, id int, opt *QueryOptions) (*SyncDataQueryResponse, error) {
	u := fmt.Sprintf("/syncs/%v/data", id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var data *SyncDataQueryResponse
	_, err = s.client.Do(ctx, req, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

