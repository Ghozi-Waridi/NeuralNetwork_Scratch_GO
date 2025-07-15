package main

import (
	// "log"
	"fmt"

)

func dot(a[][]float64, b[][]float64) [][]float64 {
	// if (len(a) != len(b)) && (len(a[1]) != len(b[1])) {
	// 	log.Fatal("ERROR!! Ukuran Matriks tidak sama")
	// 	return 0
	// }
	rowsA := len(a)
	colsA := len(a[0])
	// rowsB := len(b)
	colsB := len(b[0])

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for k := 0; k < colsB; k++ {
			sum := 0.0
			for j := 0; j < colsA; j++ {
				sum += a[i][j] * b[j][k]
			}
			result[i][k] = sum
		}
	}
	return result
}


func main() {
	a:= [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	b:= [][]float64 {
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}

	resull := dot(a, b)
	fmt.Println("Hasil Dot Product:", resull)
}












