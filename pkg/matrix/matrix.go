package matrix

import (
	"log"
)

func Dot(a [][]float64, b [][]float64) [][]float64 {
	rowsA := len(a)

	if rowsA == 0 {
		log.Println("Matrix A tidak boleh kosong")
		return nil
	}
	colsA := len(a[0])

	rowsB := len(b)
	if rowsB == 0 {
		log.Println("Matrix B tidak boleh kosong")
	}
	colsB := len(b[0])

	if colsA != rowsB {
		log.Println("Dimensi Matriks tidak valid: Ukuran antara kolom A dan baris B tidak cocok")
	}

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

func Flate(a [][]float64) []float64 {
	if len(a) == 0 {
		log.Println("Matrix tidak boleh kosong")
		return nil
	}
	result := make([]float64, 0)
	cols := len(a[0])
	rowsA := len(a)

	for i := 0; i < rowsA; i++ {
		for k := 0; k < cols; k++ {
			result = append(result, a[i][k])
		}
	}
	return result
}

func Transpose(a [][]float64) [][]float64 {
	if len(a) == 0 {
		log.Println("Matrix tidak boleh kosong")
		return nil
	}
	rowsA := len(a)
	colsA := len(a[0])

	result := make([][]float64, colsA)
	for a := range result {
		result[a] = make([]float64, rowsA)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			result[j][i] = a[i][j]
		}
	}
	return result
}
