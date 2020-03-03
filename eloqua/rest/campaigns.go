package rest

import (
	"context"
	"github.com/elqx/eloqua-go/eloqua/base"
)

type CampaignsService base.Service

type Campaign struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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
