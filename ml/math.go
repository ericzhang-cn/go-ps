package ml

import (
	"math"
)

// sigmoid computes sigmoid function value of x
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x))
}

// dot computes inner product of two one-dimension vector
func dot(vec1 map[uint64]float64, vec2 map[uint64]float64) float64 {
	var r float64
	if vec1 == nil || vec2 == nil {
		return r
	}
	for k, v1 := range vec1 {
		if v2, ok := vec2[k]; ok {
			r += v1 * v2
		}
	}
	return r
}
