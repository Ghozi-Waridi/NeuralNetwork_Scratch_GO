/*
** Package loader
Menyediakan fungsi untuk memuat gambar dari direktori dan melakukan pra-pemrosesan.

Fitur Utama:
 1. LoadImage: Memuat gambar dari path direktori.
 2. DataLoader: Mengelola pra-pemrosesan gambar (konversi ke array dan grayscale).
*/
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

/*
** LoadImage
Fungsi untuk memuat gambar dari path direktori yang diberikan.
Fungsi ini akan secara otomatis melewati file yang bukan gambar atau file gambar yang rusak.

Parameter:
  - path string: Path ke direktori yang berisi gambar.

Return:
  - []image.Image: Slice dari gambar yang berhasil dimuat.
  - error: Error jika terjadi kesalahan fatal saat membaca direktori.
*/
func LoadImage(path string) ([]image.Image, error, []string) {
	var images []image.Image
	var labels []string

	// filepath.WalkDir lebih modern dan efisien daripada filepath.Walk
	err := filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {

		if err != nil {
			return fmt.Errorf("| ERROR | Gagal mengakses path %s: %w", filePath, err)
		}

		if d.IsDir() {
			return nil
		}

		// Ini untuk menghindari error pada file non-gambar lainnya.
		fileName := strings.ToLower(d.Name())
		if !strings.HasSuffix(fileName, ".jpg") && !strings.HasSuffix(fileName, ".jpeg") && !strings.HasSuffix(fileName, ".png") {
			return nil
		}

		// Buka file gambar
		fileImage, err := os.Open(filePath)
		if err != nil {
			// Log error saat membuka file tapi jangan hentikan seluruh proses
			log.Printf("| WARNING | Gagal membuka file %s, file dilewati. Error: %v", filePath, err)
			return nil
		}
		defer fileImage.Close()

		// Dekode file menjadi data gambar
		img, _, err := image.Decode(fileImage)
		if err != nil {

			// catat sebagai peringatan dan lanjutkan ke file berikutnya, jangan hentikan semua.
			log.Printf("| WARNING | Gagal mendekode file %s, file dilewati. Error: %v", filePath, err)
			return nil
		}

		images = append(images, img)
		labels = append(labels, d.Name())
		return nil
	})

	// Error ini hanya akan terpicu jika ada masalah fatal dengan WalkDir, bukan dengan file individual.
	if err != nil {
		return nil, fmt.Errorf("| ERROR | Gagal memproses direktori secara keseluruhan: %w", err), nil
	}

	// Memberi tahu jika tidak ada gambar yang berhasil dimuat
	if len(images) == 0 {
		log.Println("| INFO | Tidak ada gambar valid yang ditemukan di path yang diberikan.")
	}

	return images, nil, labels
}

/*
** DataLoader
Fungsi untuk memuat dan mengolah data gambar.

Parameter:
  - path string: Path direktori gambar yang akan diteruskan ke LoadImage.
  - grayScale bool: Jika true, gambar akan diubah menjadi skala abu-abu.

Return:
  - [][][]uint8: Slice tiga dimensi yang berisi data gambar dalam format uint8.
    Mengembalikan nil jika terjadi kesalahan fatal atau path kosong.
*/
func DataLoader(path string, grayScale bool) ([][][]uint8, []string) {
	if path == "" {
		log.Println("| ERROR | Path tidak boleh kosong.")
		return nil, nil
	}

	img, err, labels := LoadImage(path)
	if err != nil {
		// LoadImage sudah mencatat detail error, di sini cukup log pesan umum.
		log.Printf("| FATAL | Gagal memuat dataset gambar: %v", err)
		return nil, nil
	}

	// Jika tidak ada gambar yang dimuat, tidak ada yang perlu diproses.
	if len(img) == 0 {
		return nil, nil
	}

	imagesRGB := preproces.ConvertToArray(img)

	if grayScale {
		imagesGrayscale := preproces.GrayScale(imagesRGB)
		return imagesGrayscale, labels
	}

	// Versi sebelumnya mengembalikan slice 'images' yang kosong.
	return imagesRGB, labels
}

