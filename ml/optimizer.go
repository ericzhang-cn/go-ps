package ml

// Optimizer is interface for machine learning optimizer
type Optimizer interface {
	Optimize(model map[uint64]interface{}, loss string, features []map[uint64]float64, labels []uint8) (err error)
}

// SGDOptimizer is stochastic gradient descent optimizer implementation
type SGDOptimizer struct {
}
