package model

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"neuralnetworks/pkg/Activation"
	"neuralnetworks/pkg/deactivate"
	"neuralnetworks/pkg/matrix"
	"time"
)

type HyperParameters struct {
	Architecture []int
	LearningRate float64
	Weight       [][][]float64
	Bias         [][][]float64
}

func InitHyperParameters(learningrate float64, architecture ...int) *HyperParameters {
	if len(architecture) < 2 {
		log.Fatalf("Arsitektur harus memiliki minimal 2 layer (input, output)")
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

func InitialWeights(weight [][]float64, fanIn int) {
	stdDev := math.Sqrt(2.0 / float64(fanIn))
	rows := len(weight)
	cols := len(weight[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			weight[i][j] = rand.NormFloat64() * stdDev
		}
	}
}

// FIXED: Logika forward pass disamakan dengan fungsi Train.
func (params *HyperParameters) Predict(input []float64) []float64 {
	output := matrix.FromSlice(input)

	for i := 0; i < len(params.Weight); i++ {
		// Urutan perkalian matriks yang benar: W * a + b
		output = matrix.Dot(params.Weight[i], output)
		output = matrix.Add(output, params.Bias[i])

		// Terapkan fungsi aktivasi
		if i == len(params.Weight)-1 {
			output = matrix.Map(output, activation.Sigmoid) // Aktivasi output layer
		} else {
			output = matrix.Map(output, activation.Relu) // Aktivasi hidden layer
		}
	}
	return matrix.Flate(output)
}

func CalculateMSE(errorMatrix [][]float64) float64 {
	var sum float64
	rows := len(errorMatrix)
	if rows == 0 {
		return 0.0
	}
	for i := 0; i < rows; i++ {
		err := errorMatrix[i][0]
		sum += err * err
	}
	return sum / float64(rows)
}

func (params *HyperParameters) Train(inputs [][]float64, targets [][]float64, epochs int) []float64 {
	lossHistory := make([]float64, epochs)

	for epoch := 0; epoch < epochs; epoch++ {
		totalEpochLoss := 0.0
		for i, inputSample := range inputs {
			input := matrix.FromSlice(inputSample)
			target := matrix.FromSlice(targets[i])

			// ======= FORWARD PROPAGATION =======
			activations := make([][][]float64, len(params.Architecture))
			activations[0] = input
			zs := make([][][]float64, len(params.Weight))

			for layer := 0; layer < len(params.Weight); layer++ {
				z := matrix.Dot(params.Weight[layer], activations[layer])
				z = matrix.Add(z, params.Bias[layer])
				zs[layer] = z

				var a [][]float64
				if layer == len(params.Weight)-1 {
					a = matrix.Map(z, activation.Sigmoid)
				} else {
					a = matrix.Map(z, activation.Relu)
				}
				activations[layer+1] = a
			}

			// ======= BACKWARD PROPAGATION =======
			outputError := matrix.Subtract(target, activations[len(activations)-1])
			totalEpochLoss += CalculateMSE(outputError)

			dzError := matrix.Map(zs[len(zs)-1], deactivate.Sigmoid)
			delta := matrix.Multiply(outputError, dzError)

			prevActivateT := matrix.Transpose(activations[len(activations)-2])
			dWeight := matrix.Dot(delta, prevActivateT)
			dBias := delta
			params.Weight[len(params.Weight)-1] = matrix.Add(params.Weight[len(params.Weight)-1], matrix.Map(dWeight, func(x float64) float64 { return x * params.LearningRate }))
			params.Bias[len(params.Bias)-1] = matrix.Add(params.Bias[len(params.Bias)-1], matrix.Map(dBias, func(x float64) float64 { return x * params.LearningRate }))

			for layer := len(params.Weight) - 2; layer >= 0; layer-- {
				weightsT := matrix.Transpose(params.Weight[layer+1])
				hiddenError := matrix.Dot(weightsT, delta)
				dzDeactivate := matrix.Map(zs[layer], deactivate.Relu)
				delta = matrix.Multiply(hiddenError, dzDeactivate)

				prevActivateT := matrix.Transpose(activations[layer])
				dWeight = matrix.Dot(delta, prevActivateT)
				dBias = delta
				params.Weight[layer] = matrix.Add(params.Weight[layer], matrix.Map(dWeight, func(x float64) float64 { return x * params.LearningRate }))
				params.Bias[layer] = matrix.Add(params.Bias[layer], matrix.Map(dBias, func(x float64) float64 { return x * params.LearningRate }))
			}
		}

		averageLoss := totalEpochLoss / float64(len(inputs))
		lossHistory[epoch] = averageLoss
		if (epoch+1)%100 == 0 || epoch == 0 {
			fmt.Printf("Epoch %d/%d, Loss: %f\n", epoch+1, epochs, averageLoss)
		}
	}
	return lossHistory
}

func ArgMax(slice []float64) int {
	maxIndex := -1
	macVal := -1.0
	for i, val := range slice {
		if val > macVal {
			macVal = val
			maxIndex = i
		}
	}
	return maxIndex
}

func (params *HyperParameters) Test(tesInputs, testTargets [][]float64) (accuracy float64, loss float64) {
	if len(tesInputs) != len(testTargets) {
		log.Fatalf("Jumlah input dan target tidak sesuai")
		return -1, -1
	}
	var totalLoss float64
	var correctPredictions int

	for i := 0; i < len(tesInputs); i++ {
		input := tesInputs[i]
		target := testTargets[i]
		prediction := params.Predict(input)
		predictMatrix := matrix.FromSlice(prediction)
		targetMatrix := matrix.FromSlice(target)
		errorMatrix := matrix.Subtract(predictMatrix, targetMatrix)
		totalLoss += CalculateMSE(errorMatrix)
		predictLoss := ArgMax(prediction)
		actualLoss := ArgMax(target)
		if predictLoss == actualLoss {
			correctPredictions++
		}
	}
	numSamples := float64(len(tesInputs))
	averangeLoss := totalLoss / numSamples
	accuracy = (float64(correctPredictions) / numSamples) * 100.0
	return accuracy, averangeLoss
}
