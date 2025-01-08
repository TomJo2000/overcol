package export

import (
	"cmp"
	"fmt"
	cs "github.com/TomJo2000/overcol/colorspaces"
	"strconv"
	"strings"
)

const esc_reset string = "\x1b[m"

// returns the ANSI escape sequence for the given color
func ansi_color(color cs.RGBA) string {
	const uint8_max int = 255
	var (
		luma int    = int(color.R)*299 + int(color.G)*587 + int(color.B)*114 // Luma per Rec.709
		fg   string                                                          // foreground color to use, empty meaning the terminal default
	)

	// If Luma is greater than 50% of the maximum use a black foreground color
	if luma > uint8_max*1000/2 {
		fg = "30;"
	}

	escape_seq := "\x1b[" + fg + "48;2;" + // Foreground color and 24 bit background color
		strconv.FormatUint(uint64(color.R), 10) + ";" + // Red channel
		strconv.FormatUint(uint64(color.G), 10) + ";" + // Green channel
		strconv.FormatUint(uint64(color.B), 10) + "m" // Blue channel

	return escape_seq
}

type AnsiOpts struct {
	Format  string // format string for the values
	Spacing int    // number of spaces between indicies
	Indices bool   // Print indicies instead of hex codes
	Blank   bool   // Don't print indices
}

func Ansi_Cube(cube [][][]cs.OkLAB, opts AnsiOpts) string {
	var (
		size int = len(cube)
		// size_digits uint64 = uint64(math.Floor(math.Log2(float64(opts.base)) / math.Log2(float64(size))))
		format       string = opts.Format // strings.Repeat("%"+strconv.FormatUint(size_digits, 10)+"s", 3)
		spacing      int    = opts.Spacing
		blank        bool   = opts.Blank
		use_indices  bool   = opts.Indices
		display_text string
		output       = strings.Builder{}
		segments     = make([][]string, size*size)
	)

	if blank {
		display_text = "   "
	}

	for g := 0; g < size; g++ {
		segments[g] = make([]string, size)
		for r := 0; r < size; r++ {
			line := strings.Builder{}
			for b := 0; b < size; b++ {
				var idx cs.RGBA
				if use_indices {
					idx = cs.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
				}

				var (
					color = ansi_color(cube[r][g][b].ToRGBA())
					val   = cmp.Or(idx, cube[r][g][b].ToRGBA())
					text  = cmp.Or(display_text, fmt.Sprintf(format, val.R, val.G, val.B))
				)

				// add padding between indices
				if b > 0 {
					line.WriteString(" ")
				}
				line.WriteString(color + text + esc_reset)
			}
			switch {
			case r == size/2-1 || r == size-1:
				// line.WriteString("\n")
			default:
				line.WriteString(strings.Repeat(" ", spacing))

			}
			segments[g][r] = line.String()
		}
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size/2; y++ {
			output.WriteString(segments[x][y])
		}
		output.WriteString("\n")
	}

	output.WriteString("\n")

	for x := 0; x < size; x++ {
		for y := size / 2; y < size; y++ {
			output.WriteString(segments[x][y])
		}
		output.WriteString("\n")
	}

	return output.String()
}
