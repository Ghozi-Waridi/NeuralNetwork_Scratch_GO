package deactivate

import (
	"math"
)
// pkg/deactivate/deactivate.go
func Sigmoid(x float64) float64 {
    // Hitung dulu sigmoid dari x
    s := 1 / (1 + math.Exp(-x))
    // Baru hitung turunannya
    return s * (1 - s)
}

func Relu(x float64) float64 {
	if x <= 0 {
		return 0
	}
	return 1
}
