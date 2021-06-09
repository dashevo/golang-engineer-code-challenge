package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummaryStats(t *testing.T) {
	storage := InMemory{metrics: fakeMetrics}
	aggr := storage.Aggregator()
	stats, err := aggr.SummaryStats()
	assert.NoError(t, err)
	assert.Equal(t, float64(145), stats.TotalCost)
	assert.Equal(t, int64(870), stats.TotalSize)
	assert.Equal(t, int64(70), stats.TotalElapsed)
	groupedStats := []Stats{
		{
			Name:    "GET /api/v1/get",
			Type:    "request",
			Size:    250,
			Cost:    40,
			Elapsed: 15,
		},
		{
			Name:    "p2p",
			Type:    "network",
			Size:    450,
			Cost:    75,
			Elapsed: 33,
		},
		{
			Name:    "self-hosted",
			Type:    "network",
			Size:    420,
			Cost:    70,
			Elapsed: 37,
		},
		{
			Name:    "POST /api/v1/post",
			Type:    "request",
			Size:    620,
			Cost:    105,
			Elapsed: 55,
		},
	}
	assert.Equal(t, groupedStats, stats.GroupedStats.Slice())
}
