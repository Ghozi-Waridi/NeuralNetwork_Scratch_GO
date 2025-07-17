package activation

import (
	"math"
)

func Relu(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
