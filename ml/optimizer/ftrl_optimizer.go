package optimizer

import (
	"errors"
	"math"

	"github.com/ericzhang-cn/go-ps/ml"
	"github.com/ericzhang-cn/go-ps/ml/loss"
)

// FTRLParam represents FTRL model parameters
type FTRLParam struct {
	w     float64
	z     float64
	q     float64
	theta float64
}

// FTRLOptimizer is FOLLOW-THE-REGULARIZED-LEADER optimizer implementation
type FTRLOptimizer struct {
	loss    loss.Loss
	alpha   float64
	beta    float64
	lambda1 float64
	lambda2 float64
	bias    bool
}

// Optimize update model using FOLLOW-THE-REGULARIZED-LEADER
func (o *FTRLOptimizer) Optimize(model map[uint64]interface{}, features []map[uint64]float64, labels []float64) (err error) {
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

	params := make(map[uint64]FTRLParam)
	w := make(map[uint64]float64)
	for i, v := range model {
		params[i] = v.(FTRLParam)
		w[i] = params[i].w
	}
	// gradient of loss function
	grads, err := o.loss.Gradient(features, labels, w)
	if err != nil {
		return err
	}
	// update using FTRL formula
	for i, g := range grads {
		p, ok := params[i]
		if !ok {
			params[i] = FTRLParam{}
			p = params[i]
		}
		theta := (1 / o.alpha) * (math.Sqrt(p.q+g*g) - math.Sqrt(p.q))
		p.q += g * g
		p.z += g - theta*p.w
		p.theta += theta
		if math.Abs(p.z) < o.lambda1 {
			p.w = 0
		} else {
			p.w = (1 / (-(o.lambda2 + p.theta))) * (p.z - o.lambda1*ml.Sgn(p.z))
		}
		model[i] = p
	}

	return nil
}
