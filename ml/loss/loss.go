package loss

// Loss is interface for loss function
type Loss interface {
	Loss(features []map[uint64]float64, labels []float64, weight map[uint64]float64) (loss float64, err error)
	Gradient(features []map[uint64]float64, labels []float64, weight map[uint64]float64) (grads map[uint64]float64, err error)
}
