package export

import (
	"cmp"
	"fmt"
	cs "github.com/TomJo2000/overcol/colorspaces"
	"strings"
)

type OutputFormat int

type tuple[A any, B any] struct {
	String A
	Count  B
}

type Padding tuple[string, int]
type Gaps tuple[string, int]

const (
	esc_reset string       = "\x1b[m"
	OutputHex OutputFormat = iota
	OutputIndex
	OutputValue
)

type AnsiOpts struct {
	Format  OutputFormat // format string for the values
	Padding Padding
	Gaps    Gaps
}

func Ansi_Cube(cube [][][]cs.OkLAB, opts AnsiOpts) string {
	var (
		// padding = strings.Repeat(
		// 	cmp.Or(opts.Padding.String, " "),
		// 	cmp.Or(opts.Padding.Count, 1),
		// )
		gaps = strings.Repeat(
			cmp.Or(opts.Gaps.String, " "),
			cmp.Or(opts.Gaps.Count, 3),
		)
		output       = strings.Builder{}
		size         = len(cube)
		segments     = make([][]string, size*size)
		outputFormat = func(o OutputFormat, ansiString string, val cs.RGBA) string {
			var format_string string
			switch o {
			case OutputHex:
				format_string = "%s#%02X%02X%02X%s"
			case OutputIndex:
				format_string = "%s%0d%0d%0d%s"
			case OutputValue:
				return ansiString
			}

			return fmt.Sprintf(format_string, ansiString, val.R, val.G, val.B, esc_reset)
		}

		val = func(r, g, b int) cs.RGBA { return cube[r][g][b].ToRGBA() }
	)

	if opts.Format == OutputIndex {
		val = func(r, g, b int) cs.RGBA { return cs.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0xFF} }
	}

	line := strings.Builder{}
	for g := range size {
		segments[g] = make([]string, size)
		for r := range size {
			for b := range size {
				line.WriteString(
					outputFormat(opts.Format, cube[r][g][b].ToRGBA().AnsiString(), val(r, g, b)),
				)
			}
			line.WriteString(gaps)
			segments[g][r] = line.String()
			line.Reset()
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

	for n := range size * 2 {
		switch {
		case n < size:
			fmt.Fprintln(&output, strings.Join(segments[n%size][:size/2], ""))
		case n == size:
			output.WriteString("\n")
			fallthrough
		case n >= size:
			fmt.Fprintln(&output, strings.Join(segments[n%size][size/2:], ""))
		}
	}

	return output.String()
}
