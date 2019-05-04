package parser

import (
	"encoding/csv"
	"os"
	"strconv"
)

// CsvParser is data parser for csv file format
type CsvParser struct {
	F         *os.File
	HasHeader bool
}

// Parse converts csv data to features and labels
func (p *CsvParser) Parse() (features []map[uint64]float64, labels []float64, err error) {
	reader := csv.NewReader(p.F)
	if p.HasHeader {
		_, err = reader.Read()
		if err != nil {
			return nil, nil, err
		}
	}
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	n := len(rows)

	features = make([]map[uint64]float64, n)
	labels = make([]float64, n)
	for i, row := range rows {
		features[i] = make(map[uint64]float64)
		for j, col := range row {
			if j == 0 {
				continue
			}
			if j == 1 {
				labels[i], _ = strconv.ParseFloat(col, 64)
				continue
			}
			features[i][uint64(j)], _ = strconv.ParseFloat(col, 64)
		}
	}

	return features, labels, nil
}
