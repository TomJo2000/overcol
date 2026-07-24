package colorspaces

import (
	// Used by ToLAB()
	"github.com/alltom/oklab"
	"image/color"
	// Used by AnsiString()
	"fmt"
	"math"
	"strconv"
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

// Convert RGBA value to AnsiString
func (t RGBA) AnsiString() string {
	const bg = "48;2;"
	var (
		luma = int(t.R)*299 + int(t.G)*587 + int(t.B)*114 // Luma per Rec.709
		fg   string
	)

	// If Luma is greater than 50% of the maximum use a black foreground
	if luma > math.MaxUint8*1000/2 {
		fg = "30;"
	}

	return fmt.Sprintf("\x1b[%s%s%s;%s;%sm", fg, bg,
		strconv.FormatUint(uint64(t.R), 10), // Red channel
		strconv.FormatUint(uint64(t.G), 10), // Green channel
		strconv.FormatUint(uint64(t.B), 10), // Blue channel
	)
}
