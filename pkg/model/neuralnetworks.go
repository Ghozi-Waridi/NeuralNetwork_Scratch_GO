package model

import (
	"math"
	"math/rand"
	"neuralnetworks/pkg/matrix"
	"time"
)

type HyperParameters struct {
	Architecture []int
	LearningRate float64

	Weight [][][]float64
	Bias   [][][]float64
}

func InitHyperParameters(learningrate float64, architecture ...int) *HyperParameters {
	rand.Seed(time.Now().UnixNano())

	param := &HyperParameters{
		Architecture: architecture,
		LearningRate: learningrate,
		Weight:       make([][][]float64, len(architecture)-1),
		Bias:         make([][][]float64, len(architecture)-1),
	}

	for i := 0; i < len(architecture)-1; i++ {

		rows := architecture[i+1]
		cols := architecture[i]

		param.Weight[i] = matrix.CreateMatrix(rows, cols)
		InitialWeights(param.Weight[i], cols)
		param.Bias[i] = matrix.CreateMatrix(rows, 1)
	}
	return nil
}

func InitialWeights(Weight [][]float64, fanIn int) {
	stdDev := math.Sqrt(2.0 / float64(fanIn))

	rows := len(Weight)
	cols := len(Weight[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			Weight[i][j] = rand.NormFloat64() * stdDev
		}
	}
}
