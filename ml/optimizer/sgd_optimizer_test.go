package optimizer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ericzhang-cn/go-ps/ml"

	"github.com/ericzhang-cn/go-ps/ml/loss"
	"github.com/ericzhang-cn/go-ps/ml/parser"
)

func TestSGDOptimizer(t *testing.T) {
	path, err := filepath.Abs("../../data/binary_classfication.csv")
	if err != nil {
		t.Error(err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	parser := parser.CsvParser{
		F:         file,
		HasHeader: true,
	}
	features, labels, err := parser.Parse()
	if err != nil {
		t.Error(err)
	}
	n := len(features)

	model := make(map[uint64]float64)
	w := make(map[uint64]interface{})
	loss := loss.LogisticLoss{}
	optimizer := SGDOptimizer{
		lr:      0.001,
		lambda1: 0,
		lambda2: 0,
		loss:    &loss,
		bias:    true,
	}
	batchSize := 2048
	for epoch := 1; epoch <= 2; epoch++ {
		t.Logf("====== epoch %d ======", epoch)
		ml.Shuffle(features[:n-20000], labels[:n-20000])
		for i := 0; i < n-20000; i += batchSize {
			l, err := loss.Loss(features[n-20000:], labels[n-20000:], model)
			if err != nil {
				t.Error(err)
			}
			t.Logf("iter: %d, loss: %f", i/batchSize+1, l)
			low := i
			up := i + batchSize
			if up > n-20000 {
				up = n - 20000
			}
			x := features[low:up]
			y := labels[low:up]
			err = optimizer.Optimize(w, x, y)
			for k := range w {
				model[k] = w[k].(float64)
			}
			if err != nil {
				t.Error(err)
			}
		}
	}
}
