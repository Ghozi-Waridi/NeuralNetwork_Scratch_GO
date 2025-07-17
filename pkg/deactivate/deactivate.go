package deactivate

func Sigmoid(x float64) float64 {
	return x * (1 - x)
}

func Relu(x float64) float64 {
	if x <= 0 {
		return 0
	}
	return 1
}
