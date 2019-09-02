package bulk

import (
	"context"
	"fmt"
	"net/http"
)

type ActivitiesService service

// Export represents Eloqua export definition
type Export struct {
	AreSystemTimestampsInUTC bool              `json:"areSystemTimestampsInUTC,omitempty"`
	AutoDeleteDuration       string            `json:"autoDeleteDuration,omitempty"`
	CreatedAt                string            `json:"createdAt,omitempty"`
	CreatedBy                string            `json:"createdBy,omitempty"`
	DataRetentionDuration    string            `json:"dataRetentionDuration,omitempty"`
	Name                     string            `json:"name,omitempty"`
	Fields                   map[string]string `json:"fields,omitempty"`
	Filter                   string            `json:"filter,omitempty"`
	KbUsed                   int               `json:"kbUsed,omitempty"`
	MaxRecords               int               `json:"maxRecords,omitempty"`
	UpdatedAt                string            `json:"updatedAt,omitempty"`
	UpdatedBy                string            `json:"updatedBy,omitempty"`
	Uri                      string            `json:"uri,omitempty"`
}

type ActivityField struct {
	Name string `json:"name"`
	InternalName string `json:"internalName"`
	DataType string `json:"dataType"`
	HasReadOnlyConstraint bool `json:"hasReadOnlyConstrainti,omitempty"`
	HasNotNullConstraint bool `json:"hasNotNullConstraint,omitempty"`
	HasUniquenessConstraint bool `json:"hasUniquenessConstraint,omitempty"`
	Statement string `json:"statement"`
	ActivityTypes []string `json:"activityTypes,omitempty"`
	Uri string `json:"uri"`
}

type ActivityExportSearchResponse struct {
	Count int `json:"count,omitempty"`
	HasMore bool `json:"hasMore,omitempty"`
	Items []Export `json:"items,omitempty"`
	Limit int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	TotalResults int64 `json:"totalResults,omitempty"`
}


type ActivityFieldSearchResponse struct {
	Count int `json:"count,omitempty"`
	HasMore bool `json:"hasMore,omitempty"`
	Items []ActivityField `json:"items,omitempty"`
	Limit int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	TotalResults int64 `json:"totalResults,omitempty"`
}

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-activities-exports-post.html
func (s *ActivitiesService) CreateExport(ctx context.Context, export *Export) (*Export, error) {
	req, err := s.client.NewRequest("POST", "/activities/exports", export)
	if err != nil {
		return nil, err
	}

	r :=&Export{}
	_, err = s.client.Do(ctx, req, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-activities-exports-id-delete.html
func (s *ActivitiesService) DeleteExport(ctx context.Context, id int) (*http.Response, error) {
	u := fmt.Sprintf("/activities/exports/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-activities-exports-id-data-delete.html
func (s *ActivitiesService) DeleteExportData(ctx context.Context, id int) (*http.Response, error) {
	u := fmt.Sprintf("/activities/exports/%v/data", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-activities-exports-get.html
func (s *ActivitiesService) ListExports(ctx context.Context) (*ActivityExportSearchResponse, error) {
	req, err := s.client.NewRequest("GET", "/activities/exports", nil)
	if err != nil {
		return nil, err
	}

	r := &ActivityExportSearchResponse{}
	if _, err := s.client.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAC/op-api-bulk-2.0-activities-fields-get.html
func (s *ActivitiesService) ListFields(ctx context.Context) (*ActivityFieldSearchResponse, error) {
	req, err := s.client.NewRequest("GET", "/activities/fields", nil)
	if err != nil {
		return nil, err
	}

	r := &ActivityFieldSearchResponse{}
	if _, err := s.client.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

