package rest

import (
	"context"
	"github.com/elqx/eloqua-go/eloqua/base"
)

type CampaignsService base.Service

type Campaign struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	CurrentStatus string `json:"currentStatus,omitempty"`
	Depth string `json:"depth,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UpdatedBy string `json:"updatedBy,omitempty"`
	FolderId string `json:"folderId,omitempty"`
	IsReadOnly string `json:"isReadOnly,omitempty"`
	ActualCost string `json:"actualCost,omitempty"`
	BudgetedCost string `json:"budgetedCost,omitempty"`
	CampaignCategory string `json:"campaignCategory,omitempty"`
	CampaignType string `json:"campaignType,omitempty"`
	IsEmailMarketingCampaign string `json:"isEmailMarketingCampaign,omitempty"`
	IsIncludedInROI string `json:"isIncludedInROI,omitempty"`
	IsMemberAllowedReEntry string `json:"isMemberAllowedReEntry,omitempty"`
	IsSyncedWithCRM string `json:"isSyncedWithCRM,omitempty"`
	Product string `json:"product,omitempty"`
	Region string `json:"region,omitempty"`
}

type CampaignList struct {
	Elements []Campaign `json:"elements,omitempty"`
	PageSummary
}

func (s *CampaignsService) ListCampaigns(ctx context.Context, options *GetOptions) (*CampaignList, error) {
	req, err := s.Client.NewRequest("GET", "/assets/campaigns", options, nil)
	if err != nil {
		return nil, err
	}

	r :=&CampaignList{}
	_, err = s.Client.Do(ctx, req, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}
