package model

import (
	"log"
	"math"
	"math/rand"
	"neuralnetworks/pkg/activation"
	"neuralnetworks/pkg/deactivate"
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

	if len(architecture) < 2 {
		log.Fatalf("Architecture harus memiliki minimal 2 layer (input, output)")
	}

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
	return param
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

func (params *HyperParameters) Predict(input []float64) []float64 {

	// Convert Input to matrix
	output := matrix.FromSlice(input)

	// Layer yang di hasilkan dari Loop ini sesuai dengan jumlah architecture sebab weight sudah di inisialisasi sesuai Architecture
	for i := 0; i < len(params.Weight); i++ {
		output = matrix.Dot(output, params.Weight[i])
		output = matrix.Add(output, params.Bias[i])

		if i == len(params.Weight)-1 {
			// Memberikan func activation pada setiap layer dengan func Map
			output = matrix.Map(output, activation.Sigmoid)
		} else {
			// Memberikan func activation pada setiap layer dengan func Map
			output = matrix.Map(output, activation.Relu)
		}
	}
	return matrix.Flate(output)
}

func (params *HyperParameters) Train(inputs, targets []float64) {
	input := matrix.FromSlice(inputs)
	target := matrix.FromSlice(targets)
	// Membuat slice untuk menyimpan hasil activation tetapi ini tidak berisi
	activations := make([][][]float64, len(params.Architecture))
	// Inisialisasi layer pertama dengan input
	activations[0] = input
	// Membuat slice untuk menyimpan hasil z (hasil Dari linear Regersi (dot product))
	zs := make([][][]float64, len(params.Weight))

	// ======= FORWARD PROPAGATION =======
	for i := 0; i < len(params.Weight); i++ {
		z := matrix.Dot(params.Weight[i], activations[i])
		z = matrix.Add(z, params.Bias[i])
		zs[i] = z

		var a [][]float64
		if i == len(params.Weight)-1 {
			a = matrix.Map(z, activation.Sigmoid)
		} else {
			a = matrix.Map(z, activation.Relu)
		}

		activations[i+1] = a
	}

	// ======= BACKWARD PROPAGATION =======

	// menghitung error pada output layer (loss)
	outputError := matrix.Subtract(activations[len(activations)-1], target)
	dzError := matrix.Map(zs[len(zs)-1], deactivate.Sigmoid)
	delta := matrix.Multiply(outputError, dzError)

	// gradient descent untuk weight dan Bias
	prevActivateT := matrix.Transpose(activations[len(activations)-2])
	dWeight := matrix.Dot(delta, prevActivateT)
	dBias := delta

	params.Weight[len(params.Weight)-1] = matrix.Subtract(params.Weight[len(params.Weight)-1], matrix.Map(dWeight, func(x float64) float64 { return x * params.LearningRate }))
	params.Bias[len(params.Bias)-1] = matrix.Subtract(params.Bias[len(params.Bias)-1], matrix.Map(dBias, func(x float64) float64 { return x * params.LearningRate }))

	for i := len(params.Weight) - 2; i >= 0; i-- {
		weightsT := matrix.Transpose(params.Weight[i+1])
		hiddenError := matrix.Dot(weightsT, delta)
		dzDeactivate := matrix.Map(zs[i], deactivate.Relu)
		delta = matrix.Multiply(hiddenError, dzDeactivate)

		// Menghitung Gradient Descent untuk weight dan Bias
		prevActivateT := matrix.Transpose(activations[i])
		dWeight = matrix.Dot(delta, prevActivateT)
		dBias = delta

		params.Weight[i] = matrix.Subtract(params.Weight[i], matrix.Map(dWeight, func(x float64) float64 { return x * params.LearningRate }))
		params.Bias[i] = matrix.Subtract(params.Bias[i], matrix.Map(dBias, func(x float64) float64 { return x * params.LearningRate }))

	}

}
