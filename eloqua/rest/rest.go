package rest

import (
	"log"
	"net/http"
	"net/url"

	"github.com/elqx/eloqua-go/eloqua/base"
)

const (
	baseURL = "http://login.eloqua.com/"
	version = "2.0"
)

type GetOptions struct {
	// Maximum number of entities to return.
	// Must be less than or equal to 1000 and greater than or equal to 1.
	Count int `url:"count,omitempty"`

	// Level of detail returned by the request.
	// Eloqua APIs can retrieve entities at three different levels of depth: minimal, partial, and complete.
	// Any other values passed are reset to minimal by default.
	Depth string `url:"depth,omitempty"`

	// The date and time the email was last updated.
	LastUpdatedAt int `url:"lastUpdatedAt,omitempty"`

	// Specifies the field by which list results are ordered.
	OrderBy string `url:"orderBy,omitempty"`

	// Specifies which page of entities to return (the count parameter defines the number of entities per page).
	// If the page parameter is not supplied, 1 will be used by default.
	Page int `url:"page,omitempty"`

	// Specifies the search criteria used to retrieve entities.
	Search string `url:"search,omitempty"`
}

type PageSummary struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
	Total    int `json:"total,omitempty"`
}

type RestClient struct {
	// BaseClient implements common functionality among different Eloqua API clients.
	base.BaseClient

	// Services used for talking with different parts of Eloqua API.
	Emails    *EmailsService
	Campaigns *CampaignsService
	Forms     *FormsService
}

/*
type restService struct {
	client *RestClient
}
*/

func NewClient(restURL string, httpClient *http.Client) *RestClient {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	u, err := url.Parse(restURL)
	if err != nil {
		log.Fatal(err)
	}

	bc := base.BaseClient{Client: httpClient, BaseURL: u}
	bc.Common.Client = &bc
	c := &RestClient{
		bc,
		(*EmailsService)(&bc.Common),
		(*CampaignsService)(&bc.Common),
		(*FormsService)(&bc.Common),
	}
	return c
}
