package metric

var fakeMetrics = []Metric{
	{
		Method:  "GET",
		URL:     "/api/v1/get",
		Size:    150,
		Cost:    25,
		Elapsed: 10,
		Tags: map[string]string{
			"network": "p2p",
		},
	},
	{
		Method:  "GET",
		URL:     "/api/v1/get",
		Size:    100,
		Cost:    15,
		Elapsed: 5,
		Tags: map[string]string{
			"network": "self-hosted",
		},
	},
	{
		Method:  "POST",
		URL:     "/api/v1/post",
		Size:    300,
		Cost:    50,
		Elapsed: 23,
		Tags: map[string]string{
			"network": "p2p",
		},
	},
	{
		Method:  "POST",
		URL:     "/api/v1/post",
		Size:    320,
		Cost:    55,
		Elapsed: 32,
		Tags: map[string]string{
			"network": "self-hosted",
		},
	},
}

// Metric is a metric request entry
type Metric struct {
	Method  string
	URL     string
	Size    int64
	Cost    float64
	Elapsed int64
	Tags    map[string]string
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
