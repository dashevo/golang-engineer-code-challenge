package metric

import "sync"

// List of grouped stats types
const (
	RequestStatsType = "request"
	NetworkStatsType = "network"
)

// Stats is an aggregated data grouped by specific value
type Stats struct {
	Name    string
	Type    string
	Size    int64
	Cost    float64
	Elapsed int64
}

// SummaryStats is a summary metric statistics
type SummaryStats struct {
	TotalCost    float64
	TotalElapsed int64
	TotalSize    int64
	GroupedStats OrderedStats
}

func (s *SummaryStats) addStats(name, statType string, metric Metric) {
	stats, ok := s.GroupedStats.Get(name)
	if !ok {
		s.GroupedStats.Set(name, Stats{
			Name:    name,
			Type:    statType,
			Size:    metric.Size,
			Cost:    metric.Cost,
			Elapsed: metric.Elapsed,
		})
		return
	}
	stats.Size += metric.Size
	stats.Cost += metric.Cost
	stats.Elapsed += metric.Elapsed
	s.GroupedStats.Set(name, stats)
}

// Aggregator return an aggregator to perform aggregator features on stored metrics
func (m *InMemory) Aggregator() Aggregator {
	return Aggregator{iter: m.Iter()}
}

// Aggregator is a metric aggregator component
type Aggregator struct {
	iter Iter
}

// SummaryStats is a reduce function which aggregates summary statistics by stored metrics
func (r *Aggregator) SummaryStats() (SummaryStats, error) {
	stats := SummaryStats{
		GroupedStats: NewOrderedStats(),
	}
	var metric Metric
	for r.iter.Next() {
		metric = r.iter.Value()
		stats.TotalCost += metric.Cost
		stats.TotalElapsed += metric.Elapsed
		stats.TotalSize += metric.Size
		stats.addStats(metric.Method+" "+metric.URL, RequestStatsType, metric)
		if network, ok := metric.Tags["network"]; ok {
			stats.addStats(network, NetworkStatsType, metric)
		}
	}
	return stats, nil
}

// OrderedStats is an implementation of ordered map
type OrderedStats struct {
	guard *sync.RWMutex
	keys  map[string]int
	items []Stats
}

// NewOrderedStats returns a new ordered stats component
func NewOrderedStats() OrderedStats {
	return OrderedStats{
		guard: new(sync.RWMutex),
		keys:  make(map[string]int),
	}
}

// Set sets a new stats or update existed
func (s *OrderedStats) Set(key string, stats Stats) {
	s.guard.Lock()
	defer s.guard.Unlock()
	if i, ok := s.keys[key]; ok {
		s.items[i] = stats
		return
	}
	s.items = append(s.items, stats)
	s.keys[key] = len(s.items) - 1
}

// Slice returns a copy of stored stats
func (s *OrderedStats) Slice() []Stats {
	s.guard.RLock()
	defer s.guard.RUnlock()
	var res []Stats
	return append(res, s.items...)
}

// Get returns stats data and an existence flag
func (s *OrderedStats) Get(key string) (Stats, bool) {
	s.guard.RLock()
	defer s.guard.RUnlock()
	if i, ok := s.keys[key]; ok {
		return s.items[i], true
	}
	return Stats{}, false
}
