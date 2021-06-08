package metric

import (
	"fmt"
)

// Metric is a metric request entry
type Metric struct {
	Method  string
	URL     string
	Size    int64
	Cost    float64
	Elapsed int64
}

// RequestStats is a summary data by specific request
type RequestStats struct {
	Method  string
	URL     string
	Size    int64
	Elapsed int64
}

// SummaryStats is a summary metric statistics
type SummaryStats struct {
	TotalCost    float64
	TotalElapsed int64
	TotalSize    int64
	Request      map[string]*RequestStats
}

// InMemory is a in-memory storage implementation
type InMemory struct {
	metrics []Metric
}

// NewInMemory returns a new in-memory storage
func NewInMemory(items []Metric) *InMemory {
	return &InMemory{
		metrics: items,
	}
}

// Iter returns an iterator
func (m *InMemory) Iter() Iter {
	return &sliceIter{
		metrics: m.metrics,
	}
}

// Persist puts a metric structure in a storage
func (m *InMemory) Persist(metric Metric) error {
	m.metrics = append(m.metrics, metric)
	return nil
}

// Reduce applies a reduce function to a stored data in a storage
func (m *InMemory) Reduce(reduce func(acc interface{}, metric Metric) error, acc interface{}) error {
	iter := m.Iter()
	var err error
	for iter.Next() {
		err = reduce(acc, iter.Value())
		if err != nil {
			return err
		}
	}
	return iter.Err()
}

// SummaryStatsReduce is a reduce function which aggregates summary statistics by stored metrics
func SummaryStatsReduce() func(acc interface{}, metric Metric) error {
	return func(acc interface{}, metric Metric) error {
		stats, ok := acc.(*SummaryStats)
		if !ok {
			return fmt.Errorf("given unexpected type %T", acc)
		}
		stats.TotalCost += metric.Cost
		stats.TotalElapsed += metric.Elapsed
		stats.TotalSize += metric.Size
		req, ok := stats.Request[metric.URL]
		if !ok {
			stats.Request[metric.URL] = &RequestStats{
				Method:  metric.Method,
				URL:     metric.URL,
				Size:    metric.Size,
				Elapsed: metric.Elapsed,
			}
		} else {
			req.Size += metric.Size
			req.Elapsed += metric.Elapsed
		}
		return nil
	}
}
