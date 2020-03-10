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
	IsReadOnly               string            `json:"isReadOnly"`
	ActualCost               string            `json:"actualCost"`
	BudgetedCost             string            `json:"budgetedCost"`
	CampaignCategory         string            `json:"campaignCategory"`
	CampaignType             string            `json:"campaignType"`
	IsEmailMarketingCampaign string            `json:"isEmailMarketingCampaign"`
	IsIncludedInROI          string            `json:"isIncludedInROI"`
	IsMemberAllowedReEntry   string            `json:"isMemberAllowedReEntry"`
	IsSyncedWithCRM          string            `json:"isSyncedWithCRM"`
	Product                  string            `json:"product"`
	Region                   string            `json:"region"`
	FieldValues              []FieldValue      `json:"fieldValues,omitempty"`
	BadgeId                  string            `json:"badgeId"`
	ClrEndDate               string            `json:"clrEndDate"`
	CrmId                    string            `json:"crmId"`
	EndAt                    string            `json:"endAt"`
	FirstActivation          string            `json:"firstActivation"`
	MemberCount              string            `json:"memberCount"`
	RunAsUserId              string            `json:"runAsUserId"`
	ScheduledFor             string            `json:"scheduledFor"`
	SourceTemplateId         string            `json:"sourceTemplateId"`
	StartAt                  string            `json:"startAt"`
	Elements                 []json.RawMessage `json:"elements"`
}

type FieldValue struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Name  string `json:"name"`
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
