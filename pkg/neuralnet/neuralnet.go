package neuralnet

import (
	"log"
)

func LayerFullConnect(input []float64, weight [][]float64, bias []float64) ([]float64, error) {
	numInputs := len(input)
	if numInputs == 0 {
		return nil, log.Output(1, "Input tidak boleh kosong")
	}
	numOutputs := len(bias)
	if numOutputs == 0 {
		return nil, log.Output(1, "Bias tidak boleh kosong")
	}

	outputs := make([]float64, numOutputs)

	sum := 0.0
	for i := 0; i < numOutputs; i++ {
		for j := 0; j < numInputs; j++ {
			sum += input[j] * weight[j][i]
		}
		outputs[i] = sum + bias[i]
	}

	return outputs, nil
}
