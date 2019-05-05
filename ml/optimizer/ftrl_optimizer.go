package optimizer

import "github.com/ericzhang-cn/go-ps/ml/loss"

// FTRLOptimizer is FOLLOW-THE-REGULARIZED-LEADER optimizer implementation
type FTRLOptimizer struct {
	loss    loss.Loss
	alpha   float64
	beta    float64
	lambda1 float64
	lambda2 float64
	bias    bool
}
