package ml

import (
	"testing"

	"gotest.tools/assert"
)

func TestSigmoid(t *testing.T) {
	assert.Equal(t, Sigmoid(0), 0.5)
	assert.Assert(t, Sigmoid(-20) < 1E-5)
	assert.Assert(t, 1-Sigmoid(20) < 1E-5)
}

func TestDot(t *testing.T) {
	a := map[uint64]float64{
		1: float64(1),
		3: float64(3),
		5: float64(5),
	}
	b := map[uint64]float64{
		1: float64(1),
		2: float64(2),
		3: float64(3),
		4: float64(4),
	}
	assert.Equal(t, Dot(a, b), float64(10))
}
