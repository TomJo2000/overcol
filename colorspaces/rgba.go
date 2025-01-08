package colorspaces

import (
	"github.com/alltom/oklab"
	"image/color"
)

// RGBA uses float64 for higher precision when transforming value
type RGBA color.RGBA

// RGBA.ToOkLab converts a given RGBA value to OkLAB
func (rgba_color RGBA) ToLAB() OkLAB {
	// We can't define new methods on the color.RGBA struct,
	// so convert our proxy struct to that
	// since it's what the package's conversion function wants.
	rgba := color.RGBA{
		R: rgba_color.R,
		G: rgba_color.G,
		B: rgba_color.B,
		A: rgba_color.A,
	}
	okl := oklab.OklabModel.Convert(rgba).(oklab.Oklab)

	return OkLAB{L: okl.L, A: okl.A, B: okl.B}
}
