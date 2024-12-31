package main

import (
	// "fmt"
	// "math"
	"strconv"
)

func _lerp(verts [2]OkLAB, steps int) []OkLAB {
	diff := OkLAB{
		L:     (verts[1].L - verts[0].L) / float64(steps-1),
		A:     (verts[1].A - verts[0].A) / float64(steps-1),
		B:     (verts[1].B - verts[0].B) / float64(steps-1),
		alpha: (verts[1].alpha - verts[0].alpha) / float64(steps-1),
	}

	edge := make([]OkLAB, steps)
	edge[0] = verts[0]
	edge[steps-1] = verts[1]

	for idx := 1; idx < steps; idx++ {
		edge[idx] = OkLAB{
			L:     verts[0].L + float64(idx)*diff.L,
			A:     verts[0].A + float64(idx)*diff.A,
			B:     verts[0].B + float64(idx)*diff.B,
			alpha: verts[0].alpha + float64(idx)*diff.alpha,
		}
	}
	return edge
}

func _bilerp(verts [2][2]OkLAB, steps int) [][]OkLAB {
	face := make([][]OkLAB, steps)
	for i := range face {
		face[i] = make([]OkLAB, steps)
	}

	low_col := _lerp([2]OkLAB{verts[0][0], verts[1][0]}, steps)
	high_col := _lerp([2]OkLAB{verts[0][1], verts[1][1]}, steps)

	for i := 0; i < steps; i++ {
		face[i] = _lerp([2]OkLAB{low_col[i], high_col[i]}, steps)
	}

	return face
}

func _trilerp(verts [2][2][2]OkLAB, steps int) [][][]OkLAB {
	volume := make([][][]OkLAB, steps)
	corners := make([][]OkLAB, steps)
	corners[0] = _lerp([2]OkLAB{verts[0][0][0], verts[1][0][0]}, steps)
	corners[1] = _lerp([2]OkLAB{verts[0][0][1], verts[1][0][1]}, steps)
	corners[2] = _lerp([2]OkLAB{verts[0][1][0], verts[1][1][0]}, steps)
	corners[3] = _lerp([2]OkLAB{verts[0][1][1], verts[1][1][1]}, steps)

	for i := 0; i < steps; i++ {
		plane := [2][2]OkLAB{
			{corners[0][i], corners[1][i]},
			{corners[2][i], corners[3][i]},
		}
		volume[i] = _bilerp(plane, steps)
	}

	return volume
}

func main() {
	v_rgba := [8]RGBA{
		{R: 0, G: 0, B: 0, alpha: 1.0},
		{R: 0, G: 0, B: 255, alpha: 1.0},
		{R: 0, G: 255, B: 0, alpha: 1.0},
		{R: 0, G: 255, B: 255, alpha: 1.0},
		{R: 255, G: 0, B: 0, alpha: 1.0},
		{R: 255, G: 0, B: 255, alpha: 1.0},
		{R: 255, G: 255, B: 0, alpha: 1.0},
		{R: 255, G: 255, B: 255, alpha: 1.0},
	}

	verts := [2][2][2]OkLAB{
		{
			{v_rgba[0].ToOkLab(), v_rgba[1].ToOkLab()},
			{v_rgba[2].ToOkLab(), v_rgba[3].ToOkLab()},
		}, {
			{v_rgba[4].ToOkLab(), v_rgba[5].ToOkLab()},
			{v_rgba[6].ToOkLab(), v_rgba[7].ToOkLab()},
		},
	}
	const steps = 6
	out := _trilerp(verts, steps)
	imgs := Export_Cube(out)
	for idx, img := range imgs {
		SaveImg(img, "./images/"+strconv.Itoa(idx)+".png")
	}
}

// vim: set noet ts=4 sw=4 ff=unix
