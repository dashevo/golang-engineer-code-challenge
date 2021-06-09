package metric

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/shotonoff/golang-engineer-code-challenge/internal/app/httpclient"
)

// WithMetricsMiddleware adds measure traffic middleware
func WithMetricsMiddleware(computeFn ComputeCostFunc, metrics Persister, keyvals ...string) httpclient.MiddlewareFunc {
	return func(client *http.Client) {
		tipper := &measureTrafficTipper{
			metrics: metrics,
			tags:    keyvalsToMap(keyvals),
			reqCost: computeFn,
			next:    client.Transport,
		}
		client.Transport = tipper
	}
}

type measureTrafficTipper struct {
	metrics Persister
	reqCost ComputeCostFunc
	tags    map[string]string
	next    http.RoundTripper
}

// RoundTrip ...
func (r *measureTrafficTipper) RoundTrip(req *http.Request) (*http.Response, error) {
	val, err := reqSize(req)
	if err != nil {
		return nil, err
	}
	m := Metric{
		Method: req.Method,
		URL:    req.URL.Path,
		Size:   val,
		Tags:   copyTags(r.tags),
	}
	now := time.Now()
	resp, err := r.next.RoundTrip(req)
	elapsed := time.Since(now).Milliseconds()
	if err != nil {
		return nil, err
	}
	val, err = respSize(resp)
	if err != nil {
		return nil, err
	}
	m.Size += val
	m.Elapsed = elapsed
	m.Cost = r.reqCost(m.Size, elapsed)
	err = r.metrics.Persist(m)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func respSize(resp *http.Response) (int64, error) {
	body, _ := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return int64(len(body)) + headerSize(resp.Header), nil
}

func reqSize(req *http.Request) (int64, error) {
	acc := headerSize(req.Header)
	raw, _ := req.URL.MarshalBinary()
	acc += int64(len(raw) + len(req.Method))
	if req.ContentLength > 0 {
		acc += req.ContentLength
	}
	return acc, nil
}

func headerSize(header http.Header) int64 {
	var acc int
	for h, vals := range header {
		acc += len(h)
		for _, v := range vals {
			acc += len(v)
		}
	}
	return int64(acc)
}

func keyvalsToMap(keyvals []string) map[string]string {
	m := make(map[string]string)
	var val string
	for i := 0; i < len(keyvals); i += 2 {
		val = ""
		if i < len(keyvals)-1 {
			val = keyvals[i+1]
		}
		m[keyvals[i]] = val
	}
	return m
}

func copyTags(tags map[string]string) map[string]string {
	copied := make(map[string]string)
	for k, v := range tags {
		copied[k] = v
	}
	return copied
}
