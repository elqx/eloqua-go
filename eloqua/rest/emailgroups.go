package rest

import (
	"context"
	"encoding/json"
	"github.com/elqx/eloqua-go/eloqua/base"
)

type EmailGroupsService base.Service

type EmailGroup struct {
	Id                                string          `json:"id"`
	Name                              string          `json:"name"`
	Description                       string          `json:"description,omitempty"`
	Type                              string          `json:"type"`
	CurrentStatus                     string          `json:"currentStatus"`
	Depth                             string          `json:"depth"`
	CreatedAt                         string          `json:"createdAt"`
	CreatedBy                         string          `json:"createdBy"`
	UpdatedAt                         string          `json:"updatedAt"`
	UpdatedBy                         string          `json:"updatedBy"`
	FolderId                          string          `json:"folderId,omitempty"`
	DisplayName                       string          `json:"displayName"`
	EmailFooterId                     string          `json:"emailFooterId"`
	EmailHeaderId                     string          `json:"emailHeaderId"`
	EmailIds                          []string        `json:"emailIds"`
	IsVisibleInOutlookPlugin          string          `json:"isVisibleInOutlookPlugin,omitempty"`
	IsVisibleInPublicSubscriptionList string          `json:"isVisibleInPublicSubscriptionList,omitempty"`
	Permissions                       json.RawMessage `json:"permissions,omitempty"`
	SubscriptionLandingPageId         string          `json:"subscriptionLandingPageId,omitempty"`
	SubscriptionListDataLookupId      string          `json:"subscriptionListDataLookupId,omitempty"`
	SubscriptionListId                string          `json:"subscriptionListId,omitempty"`
	UnsubscriptionLandingPageId       string          `json:"unsubscriptionLandingPageId,omitempty"`
	UnsubscriptionListDataLookupId    string          `json:"unsubscriptionListDataLookupId,omitempty"`
	UnsubscriptionListId              string          `json:"unsubscriptionListId,omitempty"`
}

type EmailGroupList struct {
	Elements []EmailGroup `json:"elements,omitempty"`
	PageSummary
}

func (s *EmailGroupsService) ListEmailGroups(ctx context.Context, options *GetOptions) (*EmailGroupList, error) {
	req, err := s.Client.NewRequest("GET", "/assets/email/groups", options, nil)
	if err != nil {
		return nil, err
	}

	r := &EmailGroupList{}
	_, err = s.Client.Do(ctx, req, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}
