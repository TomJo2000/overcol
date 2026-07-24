package export

import (
	"cmp"
	"fmt"
	"strings"
	"sync"

	cs "github.com/TomJo2000/overcol/colorspaces"
)

const (
	esc_reset string       = "\x1b[m"
	OutputHex OutputFormat = iota
	OutputIndex
	OutputValue
)

type (
	tuple[A, B any] struct {
		String A
		Count  B
	}
	OutputFormat int
	Padding      tuple[string, int]
	Gaps         tuple[string, int]
	AnsiOpts     struct {
		Format  OutputFormat
		Gaps    Gaps
		Padding Padding
	}
)

func Ansi_Cube(cube [][][]cs.OkLAB, opts AnsiOpts) string {
	var (
		size     = len(cube)
		segments = make([][]string, size)
		gaps     = strings.Repeat(
			cmp.Or(opts.Gaps.String, " "),
			cmp.Or(opts.Gaps.Count, 3),
		)
		padding = strings.Repeat(
			cmp.Or(opts.Padding.String, " "),
			cmp.Or(opts.Padding.Count, 1),
		)
		output       strings.Builder
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

	var lock sync.WaitGroup
	for g := range size {
		lock.Add(1)
		go func(g int) {
			defer lock.Done()

			segments[g] = make([]string, size)

			var line strings.Builder
			for r := range size {
				line.Reset()
				if r%(size/2) != 0 {
					line.WriteString(gaps)
				}

				for b := range size {
					if b%size != 0 {
						line.WriteString(padding)
					}
					line.WriteString(
						outputFormat(opts.Format, cube[r][g][b].ToRGBA().AnsiString(), val(r, g, b)),
					)
				}
				segments[g][r] = line.String()
			}
		}(g)
	}
	lock.Wait()

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
