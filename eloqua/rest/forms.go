package rest

import (
	"context"

	"github.com/elqx/eloqua-go/eloqua/base"
)

type FormsService base.Service

type Form struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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
