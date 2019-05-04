package ml

import (
	"math"
	"math/rand"
)

// Sigmoid computes sigmoid function value of x
func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x))
}

// Dot computes inner product of two one-dimension vector
func Dot(vec1 map[uint64]float64, vec2 map[uint64]float64) float64 {
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

// Shuffle shuffles features and labels
func Shuffle(features []map[uint64]float64, labels []float64) {
	for i := range features {
		j := rand.Intn(i + 1)
		features[i], features[j] = features[j], features[i]
		labels[i], labels[j] = labels[j], labels[i]
	}
}
