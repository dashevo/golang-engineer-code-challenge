package httpclient

import (
	"net/http"
	"time"
)

// New returns http client
func New() *http.Client {
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: http.DefaultTransport,
	}
	return client
}

// WithHeaders extends a request passed http headers
func WithHeaders(client *http.Client, values map[string]string) {
	client.Transport = &headerRoundTipper{
		values: values,
		next:   client.Transport,
	}
}

type headerRoundTipper struct {
	values map[string]string
	next   http.RoundTripper
}

// RoundTrip adds or overwrites request's user-agent header
func (r *headerRoundTipper) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range r.values {
		req.Header.Set(k, v)
	}
	return r.next.RoundTrip(req)
}
