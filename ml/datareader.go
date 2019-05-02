package ml

// DataReader is interface for dataset reader
type DataReader interface {
	Read() (sample map[uint64]float64, err error)
	ReadBatch(size uint32) (samples []map[uint64]float64, err error)
}
