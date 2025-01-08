package export

import (
	cs "github.com/TomJo2000/overcol/colorspaces"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Takes in a "plane" (2D slice of RGBA values)
// Returns an RGBA image
func Export_Plane(plane [][]cs.RGBA) image.Image {
	size := len(plane)
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			var (
				R = uint8(plane[x][y].R)
				G = uint8(plane[x][y].G)
				B = uint8(plane[x][y].B)
				A = uint8(plane[x][y].A)
			)

			color := color.RGBA{R, G, B, A}
			img.SetRGBA(y, x, color)
		}
	}
	return img
}

func Export_Cube(cube [][][]cs.OkLAB) []image.Image {
	size := len(cube)
	images := make([]image.Image, size)

	for idx, layer := range cube {
		img := image.NewRGBA(image.Rect(0, 0, size, size))
		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {

				// covert the value at X,Y to RGBA
				rgb := layer[x][y].ToRGBA()
				// set the pixels RGB value
				img.SetRGBA(y, x, color.RGBA{rgb.R, rgb.G, rgb.B, rgb.A})
			}
		}
		images[idx] = img
	}
	return images
}

// Takes in an image.Image
// and saves it as a PNG to the given filename.
// Returns error || nil
func Save_PNG(img image.Image, name string) error {
	var err error

	filename, err := os.Create(name)
	if err != nil {
		return err
	}
	defer filename.Close()
	err = png.Encode(filename, img)

	return err
}
