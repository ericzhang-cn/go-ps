package optimizer

// Optimizer is interface for machine learning optimizer
type Optimizer interface {
	Optimize(model map[uint64]interface{}, features []map[uint64]float64, labels []float64) (err error)
}
