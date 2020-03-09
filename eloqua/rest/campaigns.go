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
	FieldValues []FieldValue `json:"fieldValues,omitempty"`
	BadgeId string `json:"badgeId,omitempty"`
	ClrEndDate string `json:"clrEndDate,omitempty"`
	CrmId string `json:"crmId,omitempty"`
	EndAt string `json:"endAt,omitempty"`
	FirstActivation string `json:"firstActivation,omitempty"`
	MemberCount string `json:"memberCount,omitempty"`
	RunAsUserId string `json:"runAsUserId,omitempty"`
	ScheduledFor string `json:"scheduledFor,omitempty"`
	SourceTemplateId string `json:"sourceTemplateId,omitempty"`
	StartAt string `json:"startAt,omitempty"`
	Elements []CampaignElement `json:"elements,omitempty"`
}

type CampaignElement struct {
	Id string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	MemberCount string `json:"memberCount,omitempty"`
	MemberErrorCount string `json:"memberErrorCount,omitempty"`
	OutputTerminals []CampaignOutputTerminal `json:"outputTerminals,omitempty"`
	Position
}

type CampaignOutputTerminal struct {
	Id string `json:"id,omitempty"`
	Type string `json:"type.omitempty"`
	ConnectedId string `json:"connectedId,omitempty"`
	ConnectedType string `json:"connectedType,omitempty"`
	TerminalType string `json:"terminalType,ommitempty"`
}

type Position struct {
	Type string `json:"type,omitempty"`
	X string `json:"x,omitempty"`
	Y string `json:"y,omitempty"`
}

type FieldValue struct {
	Type string `json:"fieldValue,omitempty"`
	Id string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
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
