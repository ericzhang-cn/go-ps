package ml

import (
	"math"
)

// LossFunction is interface for loss function
type LossFunction interface {
	ComputeLoss(features []map[uint64]interface{}, labels []interface{}, weight map[uint64]interface{}) (loss float64, err error)
	Gradient(features []map[uint64]interface{}, labels []interface{}, weight map[uint64]interface{}) (grads map[uint64]float64, err error)
}

// LogisticLoss is logistic loss function implementation
// use negative log-likelihood as loss
type LogisticLoss struct{}

// ComputeLoss computes logistic loss for data and linear weights w
func (lf *LogisticLoss) ComputeLoss(features []map[uint64]interface{}, labels []interface{}, weight map[uint64]interface{}) (loss float64, err error) {
	for i, feat := range features {
		dotValue, err := dot(feat, weight)
		if err == nil {
			return float64(0), err
		}
		prob := sigmoid(dotValue)
		if labels[i] == 1 {
			loss -= math.Log(prob)
		} else {
			loss -= math.Log(1 - prob)
		}
	}
	return loss, nil
}
