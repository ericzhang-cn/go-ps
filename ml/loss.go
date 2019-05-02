package ml

import (
	"math"
)

// LossFunction is interface for loss function
type LossFunction interface {
	Loss(features []map[uint64]float64, labels []float64, weight map[uint64]float64) (loss float64, err error)
	Gradient(features []map[uint64]float64, labels []float64, weight map[uint64]float64) (grads map[uint64]float64, err error)
}

// LogisticLoss is logistic loss function implementation
// use negative log-likelihood as loss
type LogisticLoss struct{}

// Loss computes logistic loss for data and linear weights w
func (lf *LogisticLoss) Loss(features []map[uint64]float64, labels []float64, weight map[uint64]float64) (loss float64, err error) {
	n := float64(len(features))
	for i, feat := range features {
		prob := sigmoid(dot(feat, weight))
		if labels[i] == float64(1) {
			loss -= math.Log(prob)
		} else {
			loss -= math.Log(1 - prob)
		}
	}
	return loss / n, nil
}

// Gradient computes weights gradient for logistic loss function
func (lf *LogisticLoss) Gradient(features []map[uint64]float64, labels []float64, weight map[uint64]float64) (grads map[uint64]float64, err error) {
	n := float64(len(features))
	grads = make(map[uint64]float64)
	for i, feat := range features {
		prob := sigmoid(dot(feat, weight))
		for j, x := range feat {
			g := x * (prob - labels[i])
			grads[j] += g
		}
	}
	for i := range grads {
		grads[i] /= n
	}
	return grads, nil
}
