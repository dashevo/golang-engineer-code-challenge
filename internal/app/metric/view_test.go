package metric

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderSummaryStats(t *testing.T) {
	groupedStats := NewOrderedStats()
	groupedStats.Set("GET /api/v1/get", Stats{
		Name:    "GET /api/v1/get",
		Type:    "request",
		Size:    5045,
		Cost:    75.10,
		Elapsed: 521,
	})
	groupedStats.Set("POST /api/v1/post", Stats{
		Name:    "POST /api/v1/post",
		Type:    "request",
		Size:    5300,
		Cost:    75.11,
		Elapsed: 2000,
	})
	groupedStats.Set("p2p", Stats{
		Name:    "p2p",
		Type:    "network",
		Size:    5045,
		Cost:    75.10,
		Elapsed: 521,
	})
	groupedStats.Set("self-hosted", Stats{
		Name:    "self-hosted",
		Type:    "network",
		Size:    5300,
		Cost:    75.11,
		Elapsed: 2000,
	})
	stats := SummaryStats{
		TotalCost:    150.21,
		TotalElapsed: 2521,
		TotalSize:    10345,
		GroupedStats: groupedStats,
	}
	buf := bytes.Buffer{}
	err := RenderSummaryStats(&buf, stats)
	assert.NoError(t, err)
	expected := `Your total expenses: 150.2100 DASH

Grouped statistics for all performed requests
Request URL		|Size/bytes	|Elapsed/ms	|Cost/dash
GET /api/v1/get		|5045		|521		|75.100000
POST /api/v1/post	|5300		|2000		|75.110000

Grouped statistics for all used networks
Network		|Size/bytes	|Elapsed/ms	|Cost/dash
p2p		|5045		|521		|75.100000
self-hosted	|5300		|2000		|75.110000
`
	assert.Equal(t, expected, buf.String())
}
