package ml

import (
	"errors"
	"sort"
)

// Auc evalutes auc metric
func Auc(labels []float64, preds []float64) (eval float64, err error) {
	if labels == nil || preds == nil {
		return 0, errors.New("empty lables or preds")
	}
	if len(labels) != len(preds) {
		return 0, errors.New("labels and preds must have same shape")
	}
	sort.Slice(labels, func(i, j int) bool { return preds[i] > preds[j] })

	var (
		pos float64
		neg float64
	)
	for _, l := range labels {
		if l == 1 {
			pos++
		} else {
			neg++
		}
	}
	if pos == 0 || neg == 0 {
		return 0, errors.New("there must be different values of labels")
	}
	var (
		p   float64
		n   float64
		tp  float64
		fp  float64
		auc float64
	)
	for _, l := range labels {
		if l == 1 {
			p++
		} else {
			n++
		}
		auc += (n/neg - fp) * tp
		tp = p / pos
		fp = n / neg
	}

	return auc, nil
}
