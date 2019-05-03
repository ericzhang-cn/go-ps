package loss

import (
	"errors"
)

// Loss is interface for loss function
type Loss interface {
	Loss(features []map[uint64]float64, labels []float64, weight map[uint64]float64) (loss float64, err error)
	Gradient(features []map[uint64]float64, labels []float64, weight map[uint64]float64) (grads map[uint64]float64, err error)
}

// GetLossFunction return loss function implementation according to name
func GetLossFunction(loss string) (lf Loss, err error) {
	switch loss {
	case "logistic":
		return &LogisticLoss{}, nil
	case "least-square":
		return &LeastSquareLoss{}, nil
	default:
		return nil, errors.New("unknown loss function")
	}
}
