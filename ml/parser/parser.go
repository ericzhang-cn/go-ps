package parser

// Parser is interface of data parser
type Parser interface {
	Parse() (features []map[uint64]float64, labels []float64, err error)
}
