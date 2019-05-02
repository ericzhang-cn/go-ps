package ml

import (
	"github.com/ericzhang-cn/go-ps/kvstore"
)

type SGDLearner struct {
}

func (l *SGDLearner) Grad(model map[uint64]interface{}, dataReader *DataReader) (grads map[uint64]interface{}, err error) {
	return nil, nil
}

func (l *SGDLearner) Update(grads map[uint64]interface{}, store *kvstore.KvStore) (err error) {
	return nil
}
