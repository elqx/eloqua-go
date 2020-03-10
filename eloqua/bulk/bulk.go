package bulk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/elqx/eloqua-go/eloqua/pkg/auth"
	"github.com/google/go-querystring/query"
)

const (
	baseURL = "http://login.eloqua.com/"
	version = "2.0"
)

type BulkClient struct {
	baseURL *url.URL

	// HTTP client used to communicate to the API
	client *http.Client

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// Services used for talking with different parts of Eloqua API
	Activities *ActivitiesService
	Cdos       *CdosService
	Contacts   *ContactsService
	Syncs      *SyncsService
}

type service struct {
	client *BulkClient
}

type Site struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Urls struct {
	Base string `json:"base,omitempty"`
	Apis Apis   `json:"apis,omitempty"`
}

type Apis struct {
	Soap Soap `json:"soap,omitempty"`
	Rest Rest `json:"rest,omitempty"`
}

type Rest struct {
	Standard string `json:"standard,omitempty"`
	Bulk     string `json:"bulk,omitempty"`
}

type Soap struct {
	Standard       string `json:"standard,omitempty"`
	DataTransfer   string `json:"dataTransfer,omitempty"`
	Email          string `json:"email,omitempty"`
	ExternalAction string `json:"externalAction,omitempty"`
}

type AccountInfo struct {
	Site Site `json:"site,omitempty"`
	Urls Urls `json:"urls,omitempty"`
}

type Field struct {
	Name                    string `json:"name"`
	InternalName            string `json:"internalName,omitempty"`
	DataType                string `json:"dataType"`
	HasReadOnlyConstraint   bool   `json:"hasReadOnlyConstrainti"`
	HasNotNullConstraint    bool   `json:"hasNotNullConstraint"`
	HasUniquenessConstraint bool   `json:"hasUniquenessConstraint"`
	Statement               string `json:"statement"`
	Uri                     string `json:"uri,omitempty"`
	CreatedAt               string `json:"createdAt,omitempty"`
	UpdatedAt               string `json:"updatedAt,omitempty"`
}

type QueryOptions struct {
	// Specifies the maximum number of records to return.
	// If not specified, the default it 1000.
	Limit int `url:"limit,omitempty"`

	Links string `url:"links,omitempty"`

	// Specifies an offset that allows you to retrieve
	// the next batch of records. Any positive integer.
	// If not specified the default is 0.
	Offset int `url:"offset,omitempty"`

	// Specified whether Total number of sync results found should be returned.
	// If not specified, the default is true.
	TotalResults bool `url:"totalResults,omitempty"`
}

// addOptions adds the parameters in opt as URL query parameters to s.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func GetAccountInfo(username, password string) (*AccountInfo, error) {
	tr := auth.BasicAuthTransport{Username: username, Password: password}
	client := NewClient(baseURL, tr.Client())
	req, _ := client.NewRequest("GET", "/id", nil)

	ctx := context.Background()
	accountInfo := &AccountInfo{}
	_, err := client.Do(ctx, req, accountInfo)
	if err != nil {
		return nil, err
	}

	return accountInfo, nil
}

func NewClient(bulkURL string, httpClient *http.Client) *BulkClient {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	u, err := url.Parse(bulkURL)
	if err != nil {
		log.Fatal(err)
	}

	c := &BulkClient{client: httpClient, baseURL: u}
	c.common.client = c
	c.Activities = (*ActivitiesService)(&c.common)
	c.Contacts = (*ContactsService)(&c.common)
	c.Cdos = (*CdosService)(&c.common)
	c.Syncs = (*SyncsService)(&c.common)
	return c
}

func (c *BulkClient) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.baseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trainling slash, but %q does not", c.baseURL)
	}

	u, err := url.Parse(urlStr[1:])
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	x := c.baseURL.ResolveReference(u)
	req, err := http.NewRequest(method, x.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Do sends an API request and returns the API response
func (c *BulkClient) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	if err == io.EOF {
		// ignore EOF errors caused by empty response body
		err = nil
	}
	return resp, err
}
