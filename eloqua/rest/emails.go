package rest

import (
	"context"
	"github.com/elqx/eloqua-go/eloqua/base"
)

type EmailsService base.Service

type Email struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Type          string `json:"type,omitempty"`
	CurrentStatus string `json:"currentStatus,omitempty"`
	Depth         string `json:"depth,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	CreatedBy     string `json:"createdBy,omitempty"`
	UpdatedAt     string `json:"updatedAt,omitempty"`
	UpdatedBy     string `json:"updatedBy,omitempty"`
	FolderId      string `json:"folderId,omitempty"`
	Subject       string `json:"subject,omitempty"`
	Archive       string `json:"archive,omitempty"`
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

	r := &EmailList{}
	_, err = s.Client.Do(ctx, req, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}
