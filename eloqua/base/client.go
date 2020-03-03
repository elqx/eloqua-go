package base

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

type Service struct {
	Client *BaseClient
}

type BaseClient struct {
	BaseURL *url.URL
	Client *http.Client
	// Reuse a single struct instead of allocating one for each service on the heap.
	Common Service
}

func (c *BaseClient) NewRequest(method, urlStr string, opt interface{}, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trainling slash, but %q does not", c.BaseURL)
	}

	u, err := url.Parse(urlStr[1:])
	if err != nil {
		return nil, err
	}

	if err := addOptions(u, opt); err != nil {
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
	x := c.BaseURL.ResolveReference(u)
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
func (c *BaseClient) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.Client.Do(req)
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

func addOptions(u *url.URL, opt interface{}) (error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil
	}

	qs, err := query.Values(opt)
	if err != nil {
		return err
	}

	u.RawQuery = qs.Encode()
	return nil
}
