package ml

import (
	"github.com/ericzhang-cn/go-ps/kvstore"
)

// Learner is interface for machine learning
type Learner interface {
	Grad(model map[uint64]interface{}, dataReader *DataReader) (grads map[uint64]interface{}, err error)
	Update(grads map[uint64]interface{}, store *kvstore.KvStore) (err error)
}
