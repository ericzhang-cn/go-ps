package ml

import (
	"github.com/ericzhang-cn/go-ps/pkg/kvstore"
	"github.com/ericzhang-cn/go-ps/rpc"
)

// Learner is interface for machine learning models
type Learner interface {
	Update(grad *rpc.RangeKV, kv *kvstore.KvStore)
}
