package httpclient

import (
	"net/http"
	"time"
)

// MiddlewareFunc ...
type MiddlewareFunc func(client *http.Client)

// New returns http client
func New(mws ...MiddlewareFunc) *http.Client {
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: http.DefaultTransport,
	}
	withMiddleware(client, mws...)
	return client
}

// WithHeaders extends a request passed http headers
func WithHeaders(values map[string]string) MiddlewareFunc {
	return func(client *http.Client) {
		client.Transport = &headerRoundTipper{
			values: values,
			next:   client.Transport,
		}
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

func withMiddleware(client *http.Client, mws ...MiddlewareFunc) {
	for _, mw := range mws {
		mw(client)
	}
}
