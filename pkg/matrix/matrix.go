package matrix

import (
	"log"
)

func Dot(a, b [][]float64) [][]float64 {
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
		log.Fatalf("Dimensi Matriks tidak valid: Ukuran antara kolom A dan baris B tidak cocok")
	}

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for k := 0; k < colsB; k++ {
			sum := 0.0
			for j := 0; j < colsA; j++ {
				sum += a[i][j] * b[k][j]
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

func CreateMatrix(rows, cols int) [][]float64 {
	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}
	return matrix
}

func Add(a, b [][]float64) [][]float64 {
	rows := len(a)
	cols := len(a[0])

	result := CreateMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

func Multiply(a, b [][]float64) [][]float64 {
	rows := len(a)
	cols := len(a[0])

	result := CreateMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = a[i][j] * b[i][j]
		}
	}

	return result
}

func Map(m [][]float64, fn func(float64) float64) [][]float64 {
	rows := len(m)
	cols := len(m[0])
	result := CreateMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = fn(m[i][j])
		}
	}
	return result
}

func FromSlice(slice []float64) [][]float64 {
	rows := len(slice)
	m := CreateMatrix(rows, 1)
	for i := 0; i < rows; i++ {
		m[i][0] = slice[i]
	}
	return m
}

func Subtract(a, b [][]float64) [][]float64 {
	rows := len(a)
	cols := len(a[0])
	result := CreateMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = a[i][j] - b[i][j]
		}
	}
	return result
}
