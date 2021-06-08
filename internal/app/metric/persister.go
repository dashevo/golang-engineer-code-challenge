package metric

// Persister is an interface for a persisting a metric entry
type Persister interface {
	Persist(item Metric) error
}

// Iter is a metric iterator interface
type Iter interface {
	Value() Metric
	Next() bool
	Err() error
}
