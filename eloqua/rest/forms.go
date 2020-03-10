package rest

import (
	"context"

	"encoding/json"
	"github.com/elqx/eloqua-go/eloqua/base"
)

type FormsService base.Service

type Form struct {
	Id                      string            `json:"id"`
	Name                    string            `json:"name"`
	Description             string            `json:"description"`
	Type                    string            `json:"type"`
	CurrentStatus           string            `json:"currentStatus"`
	Depth                   string            `json:"depth"`
	CreatedAt               string            `json:"createdAt"`
	CreatedBy               string            `json:"createdBy"`
	UpdatedAt               string            `json:"updatedAt"`
	UpdatedBy               string            `json:"updatedBy"`
	FolderId                string            `json:"folderId,omitempty"`
	CustomCSS               string            `json:"customCSS,omitempty"`
	Archived                string            `json:"archived,omitempty"`
	DefaultKeyFieldMapping  json.RawMessage   `json:"defaultKeyFieldMapping,omitempty"`
	Elements                []json.RawMessage `json:"elements,omitempty"`
	ExternalIntegrationUrl  string            `json:"externalIntegrationUrl,omitempty"`
	FormJson                string            `json:"formJson,omitempty"`
	Html                    string            `json:"html,omitempty"`
	HtmlName                string            `json:"htmlName,omitempty"`
	IsHidden                string            `json:"isHidden,omitempty"`
	IsResponsive            string            `json:"isResponsive,omitempty"`
	Permissions             []json.RawMessage `json:"permissions,omitempty"`
	ProcessingSteps         []json.RawMessage `json:"processingSteps,omitempty"`
	ProcessingType          string            `json:"processingType,omitempty"`
	ScheduledFor            string            `json:"scheduledFor,omitempty"`
	Size                    json.RawMessage   `json:"size,omitempty"`
	SourceTemplateId        string            `json:"sourceTemplateId"`
	Style                   string            `json:"style,omitempty"`
	SubmitFailedLandingPage string            `json:"submitFailedLandingPage,omitempty"`
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

	r := &FormList{}
	_, err = s.Client.Do(ctx, req, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}
