package main

import (
	// "fmt"
	"fmt"
	"image"

	// "image/color"
	// "image/jpeg"
	_ "image/jpeg"
	_ "image/png"

	// "io/fs"
	"os"
	// "path/filepath"
)

func LoadImage(path string) (image.Image, error) {

	fileImage, errors := os.Open(path)

	if errors != nil {
		return nil, fmt.Errorf("failed to open image file %s: %w", path, errors)
	}

	defer fileImage.Close()
	image, _, err := image.Decode(fileImage)

	if err != nil {
		return nil, fmt.Errorf("failed to decode image file %s: %w", path, err)
	}

	return image, nil
}

// Convert image to pixel array
type Pixel struct {
	R, G, B uint8
}

func (p Pixel) convertToArray(img image.Image) [][]Pixel {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	pixels := make([][]Pixel, height)

	for y := 0; y < height; y++ {
		pixels[y] = make([]Pixel, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pixels[y][x] = Pixel{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
			}
		}
	}

	return pixels
}

func main() {
	image, err := LoadImage("")

	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	fmt.Println(image)

	fmt.Println("Panjang ", image.Bounds().Dx())
	fmt.Println("Lebar ", image.Bounds().Dy())

	fmt.Println(image.At(0, 0))
}
