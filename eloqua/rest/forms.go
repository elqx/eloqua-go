package rest

import (
	"context"

	"github.com/elqx/eloqua-go/eloqua/base"
)

type FormsService base.Service

type Form struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	CurrentStatus string `json:"currentStatus,omitempty"`
	Depth string `json:"depth,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UpdatedBy string `json:"updatedBy,omitempty"`
	FolderId string `json:"folderId,omitempty"`
	CustomCSS string `json:"customCSS,omitempty"`
}

type FormList struct {
	Elements []Form `json:"elements,omitempty"`
	PageSummary
}

func (s *FormsService) ListForms(ctx context.Context, options *GetOptions) (*FormList, error) {
	req, err := s.Client.NewRequest("GET", "/assets/forms", options, nil)
	if err != nil {
		return nil, err
	}

	r :=&FormList{}
	_, err = s.Client.Do(ctx, req, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}
