// package main

// import (
// 	// "fmt"
// 	"image"
// 	"image/color"
// 	_ "image/jpeg"
// 	_ "image/png"
// 	"log"
// 	"neuralnetworks/pkg/dataset"

// 	// "os"
// 	"reflect"
// 	"testing"
// )

// func TestLoadImage(t *testing.T) {
// 	t.Run("Program Berhasil mengambil gambar", func(t *testing.T) {
// 		filePath := "Gambar.png"
// 		img, err := loader.LoadImage(filePath)
// 		if err != nil {
// 			t.Fatalf("Gagal mengambil gambar: %v", err)
// 		}
// 		if img == nil {
// 			t.Fatal("Gambar Tidak ditemukan")
// 		}
// 		bounds := img.Bounds()
// 		log.Printf("Gambar berhasil diambil dengan ukuran: %dx%d", bounds.Dx(), bounds.Dy())
// 	})

// 	t.Run("Program Gagal Mengambil Gambar", func(t *testing.T) {
// 		filePath := "GambarTidakAda.png"
// 		_, err := loader.LoadImage(filePath)

// 		if err != nil {
// 			log.Printf("Gagal mengambil gambar: %v", err)
// 		} else {
// 			t.Fatal("Seharusnya gagal mengambil gambar, tetapi tidak terjadi error")
// 		}
// 		log.Println("Tes gagal mengambil gambar berhasil")
// 	})
// }

// func TestConvertToArray(t *testing.T) {
// 	testImage := image.NewRGBA(image.Rect(0, 0, 2, 1))
// 	testImage.Set(0, 0, color.RGBA{255, 0, 0, 255})
// 	testImage.Set(1, 0, color.RGBA{0, 0, 255, 255})

// 	actualArray := loader.ConvertToArray(testImage)

// 	expectedArray := [][]loader.Pixel{
// 		{
// 			{R: 255, G: 0, B: 0},
// 			{R: 0, G: 0, B: 255},
// 		},
// 	}

// 	if !reflect.DeepEqual(actualArray, expectedArray) {
// 		t.Errorf("Expected %v, but got %v", expectedArray, actualArray)
// 	}
// }

// func TestGrayScale(t *testing.T) {
// 	testImage := [][]loader.Pixel{
// 		{
// 			{R: 255, G: 50, B: 100},
// 			{R: 0, G: 200, B: 255},
// 		},
// 	}

// 	grayImage := loader.GrayScale(testImage)

// 	expectedgray := [][]int{
// 		{135, 151},
// 	}

// 	if !reflect.DeepEqual(grayImage, expectedgray) {
// 		t.Errorf("Expected %v, but got %v", expectedgray, grayImage)
// 	}
// }
