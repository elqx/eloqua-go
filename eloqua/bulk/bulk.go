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
	"strings"
	"reflect"

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
	Apis Apis `json:"apis,omitempty"`
}

type Apis struct {
	Soap Soap `json:"soap,omitempty"`
	Rest Rest `json:"rest,omitempty"`
}

type Rest struct {
	Standard string `json:"standard,omitempty"`
	Bulk string `json:"bulk,omitempty"`
}

type Soap struct {
	Standard string `json:"standard,omitempty"`
	DataTransfer string `json:"dataTransfer,omitempty"`
	Email string `json:"email,omitempty"`
	ExternalAction string `json:"externalAction,omitempty"`
}

type AccountInfo struct {
	Site Site `json:"site,omitempty"`
	Urls Urls `json:"urls,omitempty"`
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
	tr := BasicAuthTransport{Username: username, Password: password}
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

// BasicAuthTransport is an http.RoundTripper that authenticates all requests
type BasicAuthTransport struct {
	Username string // Eloqua username
	Password string // Eloqua password

	// Transport is the underlying HTTP transport to use when making requests
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface
func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// clone the request
	reqClone := new(http.Request)
	*reqClone = *req

	// deep copy req.Header
	reqClone.Header = make(http.Header, len(req.Header))
	for id, header := range req.Header {
		reqClone.Header[id] = append([]string(nil), header...)
	}

	// set authorization header
	reqClone.SetBasicAuth(t.Username, t.Password)
	return t.transport().RoundTrip(reqClone)
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authenitcation
func (t *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *BasicAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
