package metric

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/shotonoff/golang-engineer-code-challenge/internal/app/httpclient"
	"github.com/stretchr/testify/assert"
)

func TestKeyvalsToMap(t *testing.T) {
	testCases := []struct {
		keyvals  []string
		expected map[string]string
	}{
		{
			keyvals: []string{"param1", "val1", "param2", "val2"},
			expected: map[string]string{
				"param1": "val1",
				"param2": "val2",
			},
		},
		{
			keyvals: []string{"param1", "val1", "param2"},
			expected: map[string]string{
				"param1": "val1",
				"param2": "",
			},
		},
		{
			keyvals: []string{"param1"},
			expected: map[string]string{
				"param1": "",
			},
		},
		{
			keyvals:  []string{},
			expected: map[string]string{},
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, keyvalsToMap(tc.keyvals))
	}
}

func TestRequestAndResponseSize(t *testing.T) {
	const body = "plain text"
	testCases := []struct {
		req          *http.Request
		handler      func(w http.ResponseWriter, r *http.Request)
		reqExpected  int64
		respExpected int64
	}{
		{
			req: httptest.NewRequest("GET", "http://localhost", nil),
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/plain")
				w.Header().Set("Content-Length", strconv.Itoa(len(body)))
				_, _ = io.WriteString(w, body)
			},
			reqExpected:  19,
			respExpected: 48,
		},
		{
			req: httptest.NewRequest("GET", "http://localhost/api/v1/text", nil),
			handler: func(w http.ResponseWriter, r *http.Request) {
				body := "plain text"
				w.Header().Set("Content-Type", "text/plain")
				_, _ = io.WriteString(w, body)
			},
			reqExpected:  31,
			respExpected: 32,
		},
		{
			req: httptest.NewRequest("POST", "http://localhost/api/v2/post", bytes.NewBufferString("request data")),
			handler: func(w http.ResponseWriter, r *http.Request) {
				body := "plain text"
				_, _ = io.WriteString(w, body)
			},
			reqExpected:  44,
			respExpected: 47,
		},
	}
	for _, tc := range testCases {
		w := httptest.NewRecorder()
		reqActual, _ := reqSize(tc.req)
		tc.handler(w, tc.req)
		resp := w.Result()

		respActual, _ := respSize(resp)

		assert.Equal(t, tc.reqExpected, reqActual)
		assert.Equal(t, tc.respExpected, respActual)
	}
}

func TestMeasureTrafficTipper(t *testing.T) {
	const factor = 5
	var tr httpclient.RoundTripFunc = func(req *http.Request) *http.Response {
		return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString("response"))}
	}
	req := httptest.NewRequest("GET", "http://localhost/api/v1", nil)
	storage := NewInMemory(nil)
	tags := map[string]string{
		"network": "p2p",
	}
	mw := measureTrafficTipper{
		metrics: storage,
		reqCost: func(trafficSize, elapsedTime int64) float64 {
			return float64(trafficSize * factor)
		},
		tags: tags,
		next: tr,
	}
	_, err := mw.RoundTrip(req)
	assert.NoError(t, err)
	iter := storage.Iter()
	assert.True(t, iter.Next())
	m := iter.Value()
	assert.Equal(t, int64(34), m.Size)
	assert.Equal(t, "/api/v1", m.URL)
	assert.Equal(t, float64(170), m.Cost)
	assert.Equal(t, tags, m.Tags)
}
