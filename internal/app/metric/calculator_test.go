package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeP2PTrafficSize(t *testing.T) {
	testCases := []struct {
		size     int64
		elapsed  int64
		expected float64
	}{
		{
			size:     25000,
			elapsed:  1300,
			expected: 0.000263,
		},
		{
			size:     0,
			elapsed:  1300,
			expected: 0.000013,
		},
		{
			size:     25000,
			elapsed:  0,
			expected: 0.00025,
		},
		{
			size:     0,
			elapsed:  0,
			expected: 0,
		},
	}
	for _, tc := range testCases {
		assert.InDelta(t, tc.expected, ComputeP2PRequestCost(tc.size, tc.elapsed), 0.0000001)
	}
}

func TestComputeHostedTrafficSize(t *testing.T) {
	testCases := []struct {
		size     int64
		elapsed  int64
		expected float64
	}{
		{
			size:     25000,
			elapsed:  1300,
			expected: 0.025013,
		},
		{
			size:     0,
			elapsed:  1300,
			expected: 0.000013,
		},
		{
			size:     25000,
			elapsed:  0,
			expected: 0.025,
		},
		{
			size:     0,
			elapsed:  0,
			expected: 0,
		},
	}
	for _, tc := range testCases {
		assert.InDelta(t, tc.expected, ComputeHostedRequestCost(tc.size, tc.elapsed), 0.0000001)
	}
}
