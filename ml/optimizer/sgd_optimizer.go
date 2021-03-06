package optimizer

import (
	"errors"

	"github.com/ericzhang-cn/go-ps/ml/loss"
)

// SGDOptimizer is stochastic gradient descent optimizer implementation
type SGDOptimizer struct {
	loss    loss.Loss
	lr      float64
	lambda1 float64
	lambda2 float64
	bias    bool
}

// Optimize update model using stochastic gradient descent
func (o *SGDOptimizer) Optimize(model map[uint64]interface{}, features []map[uint64]float64, labels []float64) (err error) {
	if o.lr <= 0 {
		return errors.New("learning rate must be positive")
	}
	if o.lambda1 < 0 {
		return errors.New("l1 penalty factor can not be negative")
	}
	if o.lambda2 < 0 {
		return errors.New("l2 penalty factor can not be negative")
	}
	if o.loss == nil {
		return errors.New("loss function not assigned")
	}

	if o.bias {
		for i := range features {
			features[i][0] = 1
		}
	}

	w := make(map[uint64]float64)
	for i, v := range model {
		w[i] = v.(float64)
	}
	// gradient of loss function
	grads, err := o.loss.Gradient(features, labels, w)
	if err != nil {
		return err
	}
	// l1 penalty
	if o.lambda1 > 0 {
		for i := range grads {
			if w[i] > 0 {
				grads[i] += o.lambda1
			} else if w[i] < 0 {
				grads[i] -= o.lambda1
			}
		}
	}
	// l2 penalty
	if o.lambda2 > 0 {
		for i := range grads {
			grads[i] += 2 * o.lambda2 * w[i]
		}
	}
	// gradient descent
	for i := range grads {
		_, ok := model[i]
		if !ok {
			model[i] = float64(0)
		}
		model[i] = model[i].(float64) - o.lr*grads[i]
	}

	return nil
}
