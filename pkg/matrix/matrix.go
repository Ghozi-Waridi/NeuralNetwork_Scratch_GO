
package matrix

import (
	"log"
)

// Dot melakukan perkalian matriks (dot product).
func Dot(a, b [][]float64) [][]float64 {
	if len(a) == 0 || len(a[0]) == 0 {
		log.Fatal("Matriks A tidak boleh kosong untuk operasi Dot.")
	}
	if len(b) == 0 || len(b[0]) == 0 {
		log.Fatal("Matriks B tidak boleh kosong untuk operasi Dot.")
	}

	rowsA := len(a)
	colsA := len(a[0])
	rowsB := len(b)
	colsB := len(b[0])

	if colsA != rowsB {
		log.Fatalf("Dimensi Matriks tidak valid: Ukuran kolom A (%d) dan baris B (%d) tidak cocok.", colsA, rowsB)
	}

	result := CreateMatrix(rowsA, colsB)


	// Rumus untuk result[i][j] adalah jumlah dari a[i][k] * b[k][j]
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			sum := 0.0
			for k := 0; k < colsA; k++ { // k berjalan sepanjang kolom A / baris B
				sum += a[i][k] * b[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

// Flate mengubah matriks 2D menjadi slice 1D (vektor).
func Flate(a [][]float64) []float64 {
	if len(a) == 0 {
		return []float64{}
	}
	rowsA := len(a)
	colsA := len(a[0])
	result := make([]float64, 0, rowsA*colsA)

	for i := 0; i < rowsA; i++ {
		result = append(result, a[i]...)
	}
	return result
}

// Transpose mengubah baris menjadi kolom dan sebaliknya.
func Transpose(a [][]float64) [][]float64 {
	if len(a) == 0 || len(a[0]) == 0 {
		return [][]float64{}
	}
	rowsA := len(a)
	colsA := len(a[0])

	result := CreateMatrix(colsA, rowsA)

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			result[j][i] = a[i][j]
		}
	}
	return result
}

// CreateMatrix membuat matriks baru dengan ukuran yang ditentukan, diinisialisasi dengan nol.
func CreateMatrix(rows, cols int) [][]float64 {
	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}
	return matrix
}

// Add menjumlahkan dua matriks secara element-wise.
func Add(a, b [][]float64) [][]float64 {

	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		log.Fatalf("Dimensi matriks untuk operasi Add tidak cocok: (%d x %d) vs (%d x %d)", len(a), len(a[0]), len(b), len(b[0]))
	}
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

// Multiply mengalikan dua matriks secara element-wise (Hadamard product).
func Multiply(a, b [][]float64) [][]float64 {

	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		log.Fatalf("Dimensi matriks untuk operasi Multiply tidak cocok: (%d x %d) vs (%d x %d)", len(a), len(a[0]), len(b), len(b[0]))
	}
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

// Map menerapkan sebuah fungsi ke setiap elemen matriks.
func Map(m [][]float64, fn func(float64) float64) [][]float64 {
	if len(m) == 0 || len(m[0]) == 0 {
		return [][]float64{}
	}
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

// FromSlice mengubah slice 1D menjadi matriks kolom (Nx1).
func FromSlice(slice []float64) [][]float64 {
	rows := len(slice)
	m := CreateMatrix(rows, 1)
	for i := 0; i < rows; i++ {
		m[i][0] = slice[i]
	}
	return m
}

// Subtract mengurangkan dua matriks secara element-wise.
func Subtract(a, b [][]float64) [][]float64 {

	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		log.Fatalf("Dimensi matriks untuk operasi Subtract tidak cocok: (%d x %d) vs (%d x %d)", len(a), len(a[0]), len(b), len(b[0]))
	}
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

