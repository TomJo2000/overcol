package colorspaces

import (
	"github.com/alltom/oklab"
)

type OkLAB oklab.Oklab

// Converts an OkLAB value to RGBA.
func (lab_color OkLAB) ToRGBA() RGBA {
	// We can't define new methods on the package's oklab.Oklab struct,
	// so convert our proxy struct to the packages struct
	//
	// The package's RGBA() function returns 4 unit32's
	// Which may range from 0x0000 to 0xFFFF
	r, g, b, _ := oklab.Oklab{
		L: lab_color.L,
		A: lab_color.A,
		B: lab_color.B,
	}.RGBA()

	// Since these are pseudo uint16's
	// and color.RGBA takes uint8's
	// we need the upper 8 bits, so right-shift the values
	return RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}
}
