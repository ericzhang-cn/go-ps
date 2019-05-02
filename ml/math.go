package ml

import (
	"errors"
	"math"
)

// sigmoid computes sigmoid function value of x
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x))
}

// dot computes inner product of two one-dimension vector
func dot(vec1 map[uint64]interface{}, vec2 map[uint64]interface{}) (float64, error) {
	var r float64
	if vec1 == nil || vec2 == nil {
		return r, nil
	}
	for k, v1 := range vec1 {
		if v2, ok := vec2[k]; ok {
			switch v1.(type) {
			case float32:
				if _, ok = v2.(float32); ok {
					r += float64(v1.(float32) * v2.(float32))
				}
			case float64:
				if _, ok = v2.(float64); ok {
					r += v1.(float64) * v2.(float64)
				}
			default:
				return float64(0), errors.New("dot operation only support float32 or float64 type")
			}
		}
	}
	return r, nil
}
