package rest

import (
	"context"
	"encoding/json"
	"github.com/elqx/eloqua-go/eloqua/base"
)

type EmailsService base.Service

type Email struct {
	Id                  string            `json:"id"`
	Name                string            `json:"name"`
	Type                string            `json:"type"`
	CurrentStatus       string            `json:"currentStatus"`
	Depth               string            `json:"depth"`
	CreatedAt           string            `json:"createdAt"`
	CreatedBy           string            `json:"createdBy"`
	UpdatedAt           string            `json:"updatedAt"`
	UpdatedBy           string            `json:"updatedBy"`
	FolderId            string            `json:"folderId,omitempty"`
	Subject             string            `json:"subject"`
	Archive             string            `json:"archive,omitempty"`
	Attachments         []json.RawMessge  `json:"attachments,omitempty"`
	BounceBackEmail     string            `json:"bounceBackEmail,omitempty"`
	ContentSections     []json.RawMessage `json:"contentSections,omitempty"`
	Description         string            `json:"description,omitempty"`
	DynamicContents     string            `json:"dynamicContents,omitempty"`
	EmailFooterId       string            `json:"emailFooterId"`
	EmailGroupId        string            `json:"emailGroupId"`
	EmailHeaderId       string            `json:"emailHeaderId"`
	EncodingId          string            `json:"encodingid"`
	FieldMerges         []json.RawMessage `json:"fieldMerges,omitempty"`
	Files               []json.RawMessage `json:"files,omitempty"`
	FolderId            string            `json:"folderId"`
	Forms               []json.RawMessage `json:"forms,omitempty"`
	HtmlContent         json.RawMessage   `json:"htmlContent,omitempty"`
	Hyperlinks          []json.RawMessage `json:"hyperlinks,omitempty"`
	Images              []json.RawMessage `json:"image,omitempty,omitempty"`
	IsContentProtexted  string            `json:"isContentProtected,omitempty"`
	IsPlainTextEditable string            `json:"isPlainTextEditable,omitempty"`
	IsTracked           string            `json:"isTracked,omitempty"`
	Layout              string            `json:"layout,omitempty"`
	Permissions         json.RawMessage   `json:"permissions,omitempty"`
	PlainText           string            `json:"plainText,omitempty"`
	PreviewText         string            `json:"previewText,omitempty"`
	RenderMode          string            `json:"renderMode,omitempty"`
	ReplyToEmail        string            `json:"replyToEmail,omitempty"`
	ReplyToName         string            `json:"replyToName,omitempty"`
	ScheduledFor        string            `json:"scheduledFor,omitempty"`
	SenderEmail         string            `json:"senderEmail,omitempty"`
	SenderName          string            `json:"senderName,omitempty"`
	SendPlainTextOnly   string            `json:"sendPlaintextOnly,omitempty"`
	SourceTemplateId    string            `json:"sourceTemplateId,omitempty"`
	Style               string            `json:"style,omitempty"`
	VirtualMTAId        string            `json:"virtualMTAId,omitempty"`
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
