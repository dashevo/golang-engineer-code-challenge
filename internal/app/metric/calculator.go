package metric

const (
	P2PTrafficFactor    float64 = 0.00001
	HostedTrafficFactor float64 = 0.001
	ElapsedTimeFactor   float64 = 0.00001
	DashFactor          float64 = 0.001
)

// ComputeCostFunc is a function type of a calculation function
type ComputeCostFunc func(size, elapsedMs int64) float64

// ComputeP2PRequestCost returns computed cost of a request to p2p service
func ComputeP2PRequestCost(size, elapsedMs int64) float64 {
	return DashFactor * ((float64(size) * P2PTrafficFactor) + ComputeElapsedTime(elapsedMs))
}

// ComputeHostedRequestCost returns computed cost of a request to hosted service
func ComputeHostedRequestCost(size, elapsedMs int64) float64 {
	return DashFactor * ((float64(size) * HostedTrafficFactor) + ComputeElapsedTime(elapsedMs))
}

// ComputeElapsedTime returns computed cost for a request time in milliseconds
func ComputeElapsedTime(elapsedMs int64) float64 {
	return float64(elapsedMs) * ElapsedTimeFactor
}
