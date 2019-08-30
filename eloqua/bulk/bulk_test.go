package bulk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	"net/http/httptest"
	"testing"
)

func setup() (client *BulkClient, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()

	server := httptest.NewServer(mux)
	url := server.URL + "/"
	client = NewClient(url, nil)
	fmt.Printf("url %v\n", url)

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func TestNewClient(t *testing.T) {
	bulkURL := "https://bulk-url.com"
	c1 := NewClient(bulkURL, nil)

	if got, want := c1.baseURL.String(), bulkURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}

	c2 := NewClient(bulkURL, nil)
	if c1.client == c2.client {
		t.Error("NewClient return the same http.Client, but they should differ")
	}
}

func TestNewRequest(t *testing.T) {
	bulkURL := "https://bulk-url.com/"
	c := NewClient(bulkURL, nil)

	inURL, outURL := "/foo", bulkURL+"foo"
	inBody, outBody := &Export{Name: "export"}, `{"name":"export"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%v) Body is %v, want %v", inBody, got, want)
	}
}

func TestNewRequest_emptyBody(t *testing.T) {
	bulkURL := "https://bulk-url.com/"
	c := NewClient(bulkURL, nil)
	req, err := c.NewRequest("GET", ".", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if req.Body != nil {
		t.Fatalf("Constructed request contains a non-nil body")
	}
}
