package main

import (
	"fmt"
	"log"

	// "golang.org/x/text/unicode/rangetable"
)

func dot(a[][]float64, b[][]float64) [][]float64 {
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

func flate(a [][]float64) []float64 {
	if len(a) ==0 {
		log.Println("Matrix tidak boleh kosong")
		return nil
	}
	result := make([]float64, 0)
	cols :=  len(a[0])
	rowsA := len(a)


	for i := 0; i < rowsA; i++ {
		for k := 0; k < cols; k++ {
			result = append(result, a[i][k])
		}
	}
	return result
}

func transpose(a [][] float64) [][]float64 {
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
		for j := 0; j < colsA;j++ {
			result[j][i] = a[i][j]
		}
	}
	return result
}

func main() {
	a:= [][]float64{
		{6, 1, 3},
		{-1, 1, 2},
		{4, 1, 3},
	}

	b:= [][]float64 {
		{1, 5, 2},
		{-1, 0, 1},
		{3, 2, 4},
	}
	//
	// c := [][]float64 {
	// 	{1,2,3},
	// 	{4,5,6},
	// }

	d := [][]float64 {
		{1, 2, 3},
		{4, 5, 6},
	}

	transposedA := transpose(d)
	fmt.Println("Matriks A Transpose: ", transposedA)

	resull := dot(a, b)
	fmt.Println("Hasil Dot Product:", resull)
		flatResult := flate(a)
	fmt.Println("Hasil Flaten : ", flatResult)
	// resull2 := dot(a, c)
	// fmt.Print("Hasil Dot Product:", resull2)
	//
	result3 := dot(b, d)
	fmt.Println("Hasil Dot Product:", result3)
}












