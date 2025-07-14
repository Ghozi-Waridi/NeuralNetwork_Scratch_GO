/*
** Package loader
Menyediakan fungsi untuk memuat gambat dari direktori, dan preproses gambar tersebut.
func  Preprses di ambil dari package preproces

Fitur Utama:
	1. LoadImage: Memuat gambat dari path direktori.
	2. DataLoader: Mengelola gambar dari func package preproces
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
)

/*
** LoadImage
Fungsi untuk meimuat gmabar dari path direktori parameter.

Parameter:

  - Path string: path direktori gambar.

Return:

  - []image.Image: slice dari gambar yang dimuat.
  - Error: error jika terjadi kesalahan saat membaca direktori atau membuka file gambar.
*/
func LoadImage(path string) ([]image.Image, error) {
	var images []image.Image
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("| ERROR | Gagal Membaca Direktori %s: %w", path, err)
		}

		if !d.IsDir() {
			fileImage, errors := os.Open(path)

			if errors != nil {
				return fmt.Errorf("| ERROR | Gagak Membuka Gambar %s: %w", path, errors)
			}

			defer fileImage.Close()
			img, _, err := image.Decode(fileImage)

			if err != nil {
				return fmt.Errorf("| ERROR | Gagal Mendekode File : %s: %w", path, err)
			}
			images = append(images, img)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("| ERROR | Gagal Membaca File Gambar: %w", err)
	}
	return images, nil
}

/*
** DataLoader
fungsi untuk mengolah gambar.

Parameter:

  - Path string: path direktori gambar dan diteruskan ke fungsi LoadImage.
  - grayScale bool: jika true, gambar akan diubah menjadi skala abu-abu.

Return:

  - [][][]uint8: slice tiga dimensi yang berisi data gambar dalam format uint8.
*/
func DataLoader(path string, grayScale bool) [][][]uint8 {

	var images [][][]uint8

	if path == "" {
		log.Println("Path Tidak Boleh Kosong")
		return nil
	}

	img, err := LoadImage(path)

	if err != nil {
		log.Printf("Gagal Memuat Gambar: %v\n", err)
		return nil
	}

	imagesRGB := preproces.ConvertToArray(img)

	if grayScale {
		images = preproces.GrayScale(imagesRGB)
		return images
	}

	return images
}
