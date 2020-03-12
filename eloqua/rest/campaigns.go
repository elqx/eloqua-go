package rest

import (
	"context"
	"encoding/json"
	"github.com/elqx/eloqua-go/eloqua/base"
)

type CampaignsService base.Service

type Campaign struct {
	Id                       string            `json:"id"`
	Name                     string            `json:"name"`
	Description              string            `json:"description"`
	Type                     string            `json:"type"`
	CurrentStatus            string            `json:"currentStatus"`
	Depth                    string            `json:"depth"`
	CreatedAt                string            `json:"createdAt"`
	CreatedBy                string            `json:"createdBy"`
	UpdatedAt                string            `json:"updatedAt"`
	UpdatedBy                string            `json:"updatedBy"`
	FolderId                 string            `json:"folderId"`
	IsReadOnly               string            `json:"isReadOnly,omitempty"`
	ActualCost               string            `json:"actualCost,omitempty"`
	BudgetedCost             string            `json:"budgetedCost,omitempty"`
	CampaignCategory         string            `json:"campaignCategory,omitempty"`
	CampaignType             string            `json:"campaignType,omitempty"`
	IsEmailMarketingCampaign string            `json:"isEmailMarketingCampaign,omitempty"`
	IsIncludedInROI          string            `json:"isIncludedInROI,omitempty"`
	IsMemberAllowedReEntry   string            `json:"isMemberAllowedReEntry,omitempty"`
	IsSyncedWithCRM          string            `json:"isSyncedWithCRM,omitempty"`
	Product                  string            `json:"product,omitmepty"`
	Region                   string            `json:"region,omitempty"`
	FieldValues              []FieldValue      `json:"fieldValues,omitempty"`
	BadgeId                  string            `json:"badgeId,omitempty"`
	ClrEndDate               string            `json:"clrEndDate,omitempty"`
	CrmId                    string            `json:"crmId,omitempty"`
	EndAt                    string            `json:"endAt,omitempty"`
	FirstActivation          string            `json:"firstActivation,omitempty"`
	MemberCount              string            `json:"memberCount,omitempty"`
	RunAsUserId              string            `json:"runAsUserId,omitempty"`
	ScheduledFor             string            `json:"scheduledFor,omitempty"`
	SourceTemplateId         string            `json:"sourceTemplateId,omitempty"`
	StartAt                  string            `json:"startAt,omitempty"`
	Elements                 []json.RawMessage `json:"elements,omitempty"`
}

type FieldValue struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value"`
}

type CampaignList struct {
	Elements []Campaign `json:"elements"`
	PageSummary
}

func (s *CampaignsService) ListCampaigns(ctx context.Context, options *GetOptions) (*CampaignList, error) {
	req, err := s.Client.NewRequest("GET", "/assets/campaigns", options, nil)
	if err != nil {
		return nil, err
	}

	r := &CampaignList{}
	_, err = s.Client.Do(ctx, req, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}
