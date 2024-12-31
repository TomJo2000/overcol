package main

import (
	"math"
)

type OkLAB struct {
	L, A, B, alpha float64
}

func (colors OkLAB) ToRGBA() RGBA {
	l_ := colors.L + 0.3963377774*colors.A + 0.2158037573*colors.B
	m_ := colors.L - 0.1055613458*colors.A - 0.0638541728*colors.B
	s_ := colors.L - 0.0894841775*colors.A - 1.2914855480*colors.B

	l, m, s := math.Pow(l_, 3), math.Pow(m_, 3), math.Pow(s_, 3)

	ret := RGBA{
		R:     +4.0767416621*l - 3.3077115913*m + 0.2309699292*s,
		G:     -1.2684380046*l + 2.6097574011*m - 0.3413193965*s,
		B:     -0.0041960863*l - 0.7034186147*m + 1.7076147010*s,
		alpha: colors.alpha,
	}
	return ret
}
