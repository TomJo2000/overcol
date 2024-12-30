package main

import "fmt"

type RGB struct {
	R, G, B uint8
	A       float64
}

func lerp(verts [8]RGB, steps int) []RGB {

	_lerp := func(verts []RGB, step float64) RGB {
		ret := RGB{
			R: uint8(float64(verts[0].R)*(1-step) + float64(verts[1].R)),
			G: uint8(float64(verts[0].G)*(1-step) + float64(verts[1].G)),
			B: uint8(float64(verts[0].B)*(1-step) + float64(verts[1].B)),
			A: 1.0,
		}
		fmt.Printf("#%02x%02x%02x\n", ret.R, ret.G, ret.B)
		return ret
	}
	_bilerp := func(verts []RGB, x, y float64) RGB {
		return _lerp([]RGB{
			_lerp(verts[:2], x),
			_lerp(verts[2:], x),
		}, y)
	}
	_trilerp := func(verts [8]RGB, x, y, z float64) RGB {
		return _lerp([]RGB{
			_bilerp(verts[:4], x, y),
			_bilerp(verts[4:], x, y),
		}, z)
	}

	var ret = make([]RGB, steps^3)

	for r := 1; r <= steps; r++ {
		for g := 1; g <= steps; g++ {
			for b := 1; b <= steps; b++ {
				fmt.Printf("%d %d %d\n", r, g, b)
				ret = append(ret, _trilerp(verts, float64(r/steps), float64(g/steps), float64(b/steps)))
			}
		}
	}
	return ret
}

func main() {
	verts := [8]RGB{
		{R: 0, G: 0, B: 0, A: 1.0},
		{R: 0, G: 0, B: 255, A: 1.0},
		{R: 0, G: 255, B: 0, A: 1.0},
		{R: 0, G: 255, B: 255, A: 1.0},
		{R: 255, G: 0, B: 0, A: 1.0},
		{R: 255, G: 0, B: 255, A: 1.0},
		{R: 255, G: 255, B: 0, A: 1.0},
		{R: 255, G: 255, B: 255, A: 1.0},
	}
	lerp(verts, 6)
	// fmt.Println(RGB{R: 5, G: 6, B: 200, A: 1.0})
}
