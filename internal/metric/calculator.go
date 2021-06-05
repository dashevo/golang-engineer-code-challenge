package metric

const (
	P2PTrafficFactor    float64 = 0.00001
	HostedTrafficFactor float64 = 0.001
	ElapsedTimeFactor   float64 = 0.00001
	DashFactor          float64 = 0.001
)

// ComputeP2PTrafficSize returns computed cost for P2P traffic and request time
func ComputeP2PTrafficSize(trafficSize, elapsedTime int64) float64 {
	return DashFactor * ((float64(trafficSize) * P2PTrafficFactor) + ComputeElapsedTime(elapsedTime))
}

// ComputeHostedTrafficSize returns computed cost for Hosted traffic and request time
func ComputeHostedTrafficSize(trafficSize, elapsedTime int64) float64 {
	return DashFactor * ((float64(trafficSize) * HostedTrafficFactor) + ComputeElapsedTime(elapsedTime))
}

// ComputeElapsedTime returns computed cost for a request time
func ComputeElapsedTime(elapsedTime int64) float64 {
	return float64(elapsedTime) * ElapsedTimeFactor
}
