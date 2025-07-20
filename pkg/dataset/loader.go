package loader

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"log"
	"neuralnetworks/pkg/preproces"
	"os"
	"path/filepath"
	"strings"
)

// ADDED: Fungsi helper untuk meratakan data dari 3D ke 2D.
func flattenData(data3D [][][]float64) [][]float64 {
	if len(data3D) == 0 {
		return nil
	}

	numImages := len(data3D)
	// Asumsi semua gambar memiliki dimensi yang sama
	height := len(data3D[0])
	width := len(data3D[0][0])
	flattenedSize := height * width

	// Buat slice 2D dengan ukuran yang sesuai
	flattenedData := make([][]float64, numImages)

	for i := 0; i < numImages; i++ {
		// Alokasikan memori untuk setiap gambar yang diratakan
		flattenedData[i] = make([]float64, 0, flattenedSize)
		for y := 0; y < height; y++ {
			// Tambahkan seluruh baris (row) ke slice yang diratakan
			flattenedData[i] = append(flattenedData[i], data3D[i][y]...)
		}
	}

	return flattenedData
}

func LoadImage(path string) ([]image.Image, error, []string) {
	var images []image.Image
	var labels []string

	err := filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("| ERROR | Gagal mengakses path %s: %w", filePath, err)
		}
		if d.IsDir() {
			return nil
		}
		fileName := strings.ToLower(d.Name())
		if !strings.HasSuffix(fileName, ".jpg") && !strings.HasSuffix(fileName, ".jpeg") && !strings.HasSuffix(fileName, ".png") {
			return nil
		}
		fileImage, err := os.Open(filePath)
		if err != nil {
			log.Printf("| WARNING | Gagal membuka file %s, file dilewati. Error: %v", filePath, err)
			return nil
		}
		defer fileImage.Close()

		img, _, err := image.Decode(fileImage)
		if err != nil {
			log.Printf("| WARNING | Gagal mendekode file %s, file dilewati. Error: %v", filePath, err)
			return nil
		}
		images = append(images, img)
		// Ekstrak label dari nama direktori parent jika memungkinkan, atau dari nama file
		// Untuk sekarang kita tetap gunakan nama file sebagai label sementara
		labels = append(labels, d.Name())
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("| ERROR | Gagal memproses direktori secara keseluruhan: %w", err), nil
	}
	if len(images) == 0 {
		log.Println("| INFO | Tidak ada gambar valid yang ditemukan di path yang diberikan.")
	}
	return images, nil, labels
}

/*
** DataLoader
Fungsi untuk memuat dan mengolah data gambar.
Return:
 - [][]float64: Slice 2D yang berisi data gambar yang sudah diratakan (flattened).
 - []string: Slice label untuk setiap gambar.
*/
// MODIFIED: Mengubah return value menjadi [][]float64
func DataLoader(path string, grayScale bool) ([][]float64, []string) {
	if path == "" {
		log.Println("| ERROR | Path tidak boleh kosong.")
		return nil, nil
	}

	img, err, labels := LoadImage(path)
	if err != nil {
		log.Printf("| FATAL | Gagal memuat dataset gambar: %v", err)
		return nil, nil
	}

	if len(img) == 0 {
		return nil, nil
	}

	// Data gambar masih dalam format [][][]uint8 atau [][][]float64 (3D)
	var imagesData3D [][][]float64
	imagesRGB := preproces.ConvertToArray(img) // Asumsi return [][][]uint8

	if grayScale {
		imagesData3D = preproces.GrayScale(imagesRGB) // Asumsi return [][][]float64
	} else {
		// Jika tidak grayscale, perlu konversi dari uint8 ke float64
		// (Implementasi ini mungkin perlu Anda sesuaikan berdasarkan paket `preproces`)
		imagesData3D = make([][][]float64, len(imagesRGB))
		for i, img3D := range imagesRGB {
			imagesData3D[i] = make([][]float64, len(img3D))
			for j, row := range img3D {
				imagesData3D[i][j] = make([]float64, len(row))
				for k, pixel := range row {
					imagesData3D[i][j][k] = float64(pixel) / 255.0 // Normalisasi
				}
			}
		}
	}

	// Ratakan data dari 3D menjadi 2D sebelum dikembalikan
	flattenedImages := flattenData(imagesData3D)

	return flattenedImages, labels
}
