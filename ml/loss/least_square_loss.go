package loss

import (
	"math"

	"github.com/ericzhang-cn/go-ps/ml"
)

// LeastSquareLoss is least square loss function implementation
type LeastSquareLoss struct{}

// Loss computes least square loss of data and model
func (lf *LeastSquareLoss) Loss(features []map[uint64]float64, labels []float64, weight map[uint64]float64) (loss float64, err error) {
	n := float64(len(features))
	for i, feat := range features {
		pred := ml.Dot(feat, weight)
		loss += math.Pow((pred - labels[i]), 2)
	}
	return loss / n, nil
}

// Gradient computes weights gradient for least square loss function
func (lf *LeastSquareLoss) Gradient(features []map[uint64]float64, labels []float64, weight map[uint64]float64) (grads map[uint64]float64, err error) {
	n := float64(len(features))
	grads = make(map[uint64]float64)
	for i, feat := range features {
		pred := ml.Dot(feat, weight)
		for j, x := range feat {
			g := x * (pred - labels[i])
			grads[j] += g
		}
	}
	for i := range grads {
		grads[i] /= n
	}
	return grads, nil
}
