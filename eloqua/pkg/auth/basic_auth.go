package auth

import (
	"net/http"
)

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
