package httpclient

import (
	"net/http"
	"time"

	"github.com/shotonoff/golang-engineer-code-challenge/internal/app/metric"
)

// WithMetricsMiddleware adds measure traffic middleware
func WithMetricsMiddleware(computeFn metric.ComputeCostFunc, metrics metric.Persister, tags ...string) MiddlewareFunc {
	return func(client *http.Client) {
		tipper := &measureTrafficTipper{
			metrics: metrics,
			tags:    tags,
			reqCost: computeFn,
			next:    client.Transport,
		}
		client.Transport = tipper
	}
}

type measureTrafficTipper struct {
	metrics metric.Persister
	reqCost metric.ComputeCostFunc
	tags    []string
	next    http.RoundTripper
}

// RoundTrip ...
func (r *measureTrafficTipper) RoundTrip(req *http.Request) (*http.Response, error) {
	val, err := calcReqSize(req)
	if err != nil {
		return nil, err
	}
	m := metric.Metric{
		Method: req.Method,
		URL:    req.URL.Path,
		Size:   val,
	}
	now := time.Now()
	resp, err := r.next.RoundTrip(req)
	elapsed := int64(time.Now().Sub(now))
	if err != nil {
		return nil, err
	}
	val, err = calcRespSize(resp)
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

func calcRespSize(resp *http.Response) (int64, error) {
	var acc int64
	acc += resp.ContentLength
	return acc, nil
}

func calcReqSize(req *http.Request) (int64, error) {
	var acc int64
	raw, _ := req.URL.MarshalBinary()
	acc += int64(len(raw) + len(req.Method))
	if req.Method != "HEAD" && req.Method != "GET" {
		acc += req.ContentLength
	}
	return acc, nil
}
