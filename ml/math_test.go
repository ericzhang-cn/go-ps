package ml

import (
	"testing"

	"gotest.tools/assert"
)

func TestSigmoid(t *testing.T) {
	assert.Equal(t, sigmoid(0), 0.5)
	assert.Assert(t, sigmoid(-20) < 1E-5)
	assert.Assert(t, 1-sigmoid(20) < 1E-5)
}

func TestDot(t *testing.T) {
	a := map[uint64]interface{}{
		1: float32(1),
		3: float32(3),
		5: float32(5),
	}
	b := map[uint64]interface{}{
		1: float32(1),
		2: float32(2),
		3: float32(3),
		4: float32(4),
	}
	r, err := dot(a, b)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, r, float64(10))
}
