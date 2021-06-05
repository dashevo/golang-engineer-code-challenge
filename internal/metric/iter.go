package metric

type sliceIter struct {
	currPos int
	metrics []Metric
	val     Metric
	err     error
}

// Value returns a metric entry
func (i *sliceIter) Value() Metric {
	return i.val
}

// Next returns a true if a value read successful otherwise false
func (i *sliceIter) Next() bool {
	if i.currPos >= len(i.metrics) {
		return false
	}
	i.val = i.metrics[i.currPos]
	i.currPos++
	return true
}

// Err returns an error if occurred during iterating
func (i *sliceIter) Err() error {
	return i.err
}
