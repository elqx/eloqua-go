package rest

import (
	"context"
	"github.com/elqx/eloqua-go/eloqua/base"
)

type EmailsService base.Service

type Email struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EmailList struct {
	Elements []Email `json:"elements,omitempty"`
	PageSummary
}

func (s *EmailsService) ListEmails(ctx context.Context, options *GetOptions) (*EmailList, error) {
	req, err := s.Client.NewRequest("GET", "/assets/emails", options, nil)
	if err != nil {
		return nil, err
	}

	r :=&EmailList{}
	_, err = s.Client.Do(ctx, req, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}
