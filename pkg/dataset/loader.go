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

func LoadImage(path string) ([]image.Image, error) {
	// Func ini masih salah logika, seharusnya func bisa mengatasi file gambar yang banyak dari sebauh folder bukan hanya 1 gambar saja.

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
