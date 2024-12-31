package main

import (
	"math"
)

type RGBA struct {
	R, G, B, alpha float64
}

func (colors RGBA) ToOkLab() OkLAB {
	l := 0.4122214708*colors.R + 0.5363325363*colors.G + 0.0514459929*colors.B
	m := 0.2119034982*colors.R + 0.6806995451*colors.G + 0.1073969566*colors.B
	s := 0.0883024619*colors.R + 0.2817188376*colors.G + 0.6299787005*colors.B

	l_, m_, s_ := math.Pow(l, 1/3), math.Pow(m, 1/3), math.Pow(s, 1/3)

	return OkLAB{
		L:     0.2104542553*l_ + 0.7936177850*m_ - 0.0040720468*s_,
		A:     1.9779984951*l_ - 2.4285922050*m_ + 0.4505937099*s_,
		B:     0.0259040371*l_ + 0.7827717662*m_ - 0.8086757660*s_,
		alpha: colors.alpha,
	}
}
