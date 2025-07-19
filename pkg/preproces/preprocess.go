/*
** preproces
Menyediakan fungsi untuk mengonversi gambar menjadi array(slice) dan pengolahan gambar.

Fitur Utama:
	1. ConvertToArray: Mengonversi gambar menjadi array pixel.
	2. GrayScale: Mengubah gambar menjadi skala abu-abu.

*/

package preproces

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
)

type Pixel struct {
	R, G, B uint8
}

/*
** ConvertToArray
Fungsi untuk mengonversi gambar menjadi array struct Pixel.

Parameter:
  - img []image.Image: slice dari gambar yang akan dikonversi.

Return:
  - [][][]Pixel: slice tiga dimensi yang berisi pixel RGB dari gambar.
*/
func ConvertToArray(images []image.Image) [][][]uint8 {
	// Slice untuk menampung semua data gambar yang sudah dikonversi
	var dataset [][][]uint8

	// Looping untuk setiap gambar di dalam slice
	for _, img := range images {
		bounds := img.Bounds()
		width, height := bounds.Dx(), bounds.Dy()

		// Slice untuk menampung data satu gambar ([baris][piksel])
		var singleImage [][]uint8

		// Iterasi per baris (koordinat y)
		for y := 0; y < height; y++ {
			// Slice untuk menampung semua piksel dalam satu baris
			var rowPixels []uint8

			// Iterasi per piksel dalam satu baris (koordinat x)
			for x := 0; x < width; x++ {
				// Ambil warna pada piksel (x, y)
				// Perhatikan: img.At() menggunakan koordinat absolut dari bounds
				r, g, b, _ := img.At(x+bounds.Min.X, y+bounds.Min.Y).RGBA()

				// Konversi channel warna dari format uint32 (0-65535) ke uint8 (0-255)
				// Ini dilakukan dengan menggeser bit ke kanan sebanyak 8 (sama dengan dibagi 257)
				pixelR := uint8(r >> 8)
				pixelG := uint8(g >> 8)
				pixelB := uint8(b >> 8)

				// Tambahkan data R, G, B ke slice baris
				rowPixels = append(rowPixels, pixelR, pixelG, pixelB)
			}
			// Tambahkan slice baris yang sudah terisi ke dalam slice gambar tunggal
			singleImage = append(singleImage, rowPixels)
		}
		// Tambahkan data gambar tunggal ke dalam dataset
		dataset = append(dataset, singleImage)
	}

	return dataset
}

/*
** GrayScale
Fungsi untuk mengubah gambar menjadi skala abu-abu.

Parameter:
  - img [][][]Pixel: slice tiga dimensi yang berisi pixel RGB dari gambar.

Return:
  - [][][]uint8: slice tiga dimensi yang berisi nilai skala abu-abu dari gambar.
*/
func GrayScale(imagesRGB [][][]uint8) [][][]uint8 {
	// Slice untuk menampung seluruh dataset yang sudah di-grayscale
	var datasetGray [][][]uint8

	// Iterasi untuk setiap gambar dalam dataset
	for _, singleImageRGB := range imagesRGB {
		// Slice untuk menampung satu gambar yang sudah di-grayscale
		var singleImageGray [][]uint8

		// Iterasi untuk setiap baris piksel dalam satu gambar
		for _, rowRGB := range singleImageRGB {
			// Slice untuk menampung satu baris yang sudah di-grayscale
			// Panjangnya adalah 1/3 dari baris RGB karena setiap 3 byte (RGB) menjadi 1 byte (Gray)
			newRowGray := make([]uint8, 0, len(rowRGB)/3)

			// Iterasi pada baris RGB, melompat 3 byte setiap kali (untuk setiap piksel)
			for i := 0; i < len(rowRGB); i += 3 {
				// Ambil nilai R, G, B
				r := rowRGB[i]
				g := rowRGB[i+1]
				b := rowRGB[i+2]

				// Hitung nilai gray menggunakan metode rata-rata (Average Method)
				// PENTING: Konversi ke uint16 dulu untuk mencegah overflow saat penjumlahan
				// (contoh: 255 + 255 + 255 akan melebihi batas maks uint8 yaitu 255)
				grayValue := (uint16(r) + uint16(g) + uint16(b)) / 3

				// Tambahkan nilai gray yang sudah dihitung (kembalikan ke uint8) ke baris baru
				newRowGray = append(newRowGray, uint8(grayValue))
			}
			// Tambahkan baris grayscale ke gambar grayscale
			singleImageGray = append(singleImageGray, newRowGray)
		}
		// Tambahkan gambar grayscale ke dataset grayscale
		datasetGray = append(datasetGray, singleImageGray)
	}

	return datasetGray
}
