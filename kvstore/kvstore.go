package kvstore

// KvStore is interface for key-value store
type KvStore interface {
	Get(key uint64) (value []byte, err error)
	Put(key uint64, value []byte) (err error)
	GetBatch(keys []uint64) (values [][]byte, err error)
	PutBatch(keys []uint64, values [][]byte) (err error)
	GetRange(begin uint64, end uint64) (values [][]byte, err error)
}
