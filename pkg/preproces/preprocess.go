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
func ConvertToArray(img []image.Image) [][][]Pixel {

	panjang := len(img)
	var imagesRGB [][][]Pixel

	for i := 0; i < panjang; i++ {
		width, height := img[i].Bounds().Dx(), img[i].Bounds().Dy()
		pixels := make([][]Pixel, height)

		for y := 0; y < height; y++ {
			pixels[y] = make([]Pixel, width)
			for x := 0; x < width; x++ {
				r, g, b, _ := img[i].At(x, y).RGBA()
				pixels[y][x] = Pixel{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
				}
			}
		}
		imagesRGB = append(imagesRGB, pixels)
	}
	return imagesRGB
}

/*
** GrayScale
Fungsi untuk mengubah gambar menjadi skala abu-abu.

Parameter:
  - img [][][]Pixel: slice tiga dimensi yang berisi pixel RGB dari gambar.

Return:
  - [][][]uint8: slice tiga dimensi yang berisi nilai skala abu-abu dari gambar.
*/
func GrayScale(img [][][]Pixel) [][][]uint8 {
	panjang := len(img)
	var datagrayScale [][][]uint8

	for i := 0; i < panjang; i++ {
		width, height := len(img[0]), len(img)

		grayPixels := make([][]uint8, height)

		for y := 0; y < height; y++ {
			grayPixels[y] = make([]uint8, width)
			for x := 0; x < width; x++ {
				gray := (uint8(img[i][y][x].R) +
					uint8(img[i][y][x].G) +
					uint8(img[i][y][x].B)) / 3

				grayPixels[y][x] = gray
			}
		}
		datagrayScale = append(datagrayScale, grayPixels)
	}
	return datagrayScale
}
