package loader

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
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

func ConvertToArray(img image.Image) [][]Pixel {
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

func GrayScale(img [][]Pixel) [][]int {
	width, height := len(img[0]), len(img)

	grayPixels := make([][]int, height)

	for y := 0; y < height; y++ {
		grayPixels[y] = make([]int, width)
		for x := 0; x < width; x++ {
			gray := (int(img[y][x].R) +
				int(img[y][x].G) +
				int(img[y][x].B)) / 3

			grayPixels[y][x] = gray
		}
	}
	return grayPixels
}
