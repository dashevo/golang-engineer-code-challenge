package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemory(t *testing.T) {
	storage := InMemory{metrics: nil}
	for _, m := range fakeMetrics {
		err := storage.Persist(m)
		assert.NoError(t, err)
	}
	var actual []Metric
	iter := storage.Iter()
	for iter.Next() {
		actual = append(actual, iter.Value())
	}
	assert.NoError(t, iter.Err())
	assert.Equal(t, fakeMetrics, actual)
}
