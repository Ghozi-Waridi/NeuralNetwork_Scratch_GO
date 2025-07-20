package encoder

import (
	"log"

	"neuralnetworks/pkg/model"
)

type Encoder struct {
	labalToindex map[string]int
	indecTolabel []string
	numClasses   int
}

func NewEncoder(labels []string) *Encoder {
	return &Encoder{
		labalToindex: make(map[string]int),
		indecTolabel: []string{},
	}
}

func (e *Encoder) Fit(labels []string) {
	e.labalToindex = make(map[string]int)
	e.indecTolabel = []string{}
	uniqueLabels := make(map[string]bool)
	for _, label := range labels {
		if !uniqueLabels[label] {
			uniqueLabels[label] = true
			e.indecTolabel = append(e.indecTolabel, label)
		}
	}
	for i, label := range e.indecTolabel {
		e.labalToindex[label] = i
	}
	e.numClasses = len(e.indecTolabel)
}

func (e *Encoder) Transform(labels []string) [][]float64 {
	oneHotEncoded := make([][]float64, len(labels))
	for i, label := range labels {
		onHotVector := make([]float64, e.numClasses)
		index, ok := e.labalToindex[label]
		if !ok {
			log.Printf("Peringatan: Label '%s' tidak ditemukan.", label)
			continue
		}
		onHotVector[index] = 0.1
		oneHotEncoded[i] = onHotVector
	}
	return oneHotEncoded
}

func (e *Encoder) InverseTransform(prediction []float64) string {
	if len(prediction) != e.numClasses {
		return "| ERROR | Ukuran Prediksi tidak sesuai dengan jumlah kelas."
	}
	maxIndex := model.ArgMax(prediction)
	if maxIndex < 0 || maxIndex >= len(e.indecTolabel) {
		return "| ERROR | Indeks prediksi di luar jangkauan."
	}
	return e.indecTolabel[maxIndex]
}
