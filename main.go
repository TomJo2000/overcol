package main

import (
	"fmt"

	cs "github.com/TomJo2000/overcol/colorspaces"
	"github.com/TomJo2000/overcol/export"
)

func _lerp(verts [2]cs.OkLAB, steps int) []cs.OkLAB {
	diff := cs.OkLAB{
		L: (verts[1].L - verts[0].L) / float64(steps-1),
		A: (verts[1].A - verts[0].A) / float64(steps-1),
		B: (verts[1].B - verts[0].B) / float64(steps-1),
		// alpha: (verts[1].alpha - verts[0].alpha) / float64(steps-1),
	}

	edge := make([]cs.OkLAB, steps)
	edge[0] = verts[0]
	edge[steps-1] = verts[1]

	for idx := 1; idx < steps; idx++ {
		edge[idx] = cs.OkLAB{
			L: verts[0].L + float64(idx)*diff.L,
			A: verts[0].A + float64(idx)*diff.A,
			B: verts[0].B + float64(idx)*diff.B,
			// alpha: verts[0].alpha + float64(idx)*diff.alpha,
		}
	}
	return edge
}

func _bilerp(verts [2][2]cs.OkLAB, steps int) [][]cs.OkLAB {
	face := make([][]cs.OkLAB, steps)
	for i := range face {
		face[i] = make([]cs.OkLAB, steps)
	}

	low_col := _lerp([2]cs.OkLAB{verts[0][0], verts[1][0]}, steps)
	high_col := _lerp([2]cs.OkLAB{verts[0][1], verts[1][1]}, steps)

	for i := 0; i < steps; i++ {
		face[i] = _lerp([2]cs.OkLAB{low_col[i], high_col[i]}, steps)
	}

	return face
}

func _trilerp(verts [2][2][2]cs.OkLAB, steps int) [][][]cs.OkLAB {
	volume := make([][][]cs.OkLAB, steps)
	corners := make([][]cs.OkLAB, steps)
	corners[0] = _lerp([2]cs.OkLAB{verts[0][0][0], verts[1][0][0]}, steps)
	corners[1] = _lerp([2]cs.OkLAB{verts[0][0][1], verts[1][0][1]}, steps)
	corners[2] = _lerp([2]cs.OkLAB{verts[0][1][0], verts[1][1][0]}, steps)
	corners[3] = _lerp([2]cs.OkLAB{verts[0][1][1], verts[1][1][1]}, steps)

	for i := 0; i < steps; i++ {
		plane := [2][2]cs.OkLAB{
			{corners[0][i], corners[1][i]},
			{corners[2][i], corners[3][i]},
		}
		volume[i] = _bilerp(plane, steps)
	}

	return volume
}

// func rgba_min_max(hex_colors []string) {
// 	return
// }

func main() {
	// start := time.Now()
	lo, hi := cs.RGBA{R: 0x05, G: 0x03, B: 0x04}, cs.RGBA{R: 0xFB, G: 0xF6, B: 0xFD, A: 0xFF}
	v_rgba := [8]cs.RGBA{
		{R: lo.R, G: lo.G, B: lo.B, A: hi.A},
		{R: lo.R, G: lo.G, B: hi.B, A: hi.A},
		{R: lo.R, G: hi.G, B: lo.B, A: hi.A},
		{R: lo.R, G: hi.G, B: hi.B, A: hi.A},
		{R: hi.R, G: lo.G, B: lo.B, A: hi.A},
		{R: hi.R, G: lo.G, B: hi.B, A: hi.A},
		{R: hi.R, G: hi.G, B: lo.B, A: hi.A},
		{R: hi.R, G: hi.G, B: hi.B, A: hi.A},
	}
	verts := [2][2][2]cs.OkLAB{
		{
			{v_rgba[0].ToLAB(), v_rgba[1].ToLAB()},
			{v_rgba[2].ToLAB(), v_rgba[3].ToLAB()},
		}, {
			{v_rgba[4].ToLAB(), v_rgba[5].ToLAB()},
			{v_rgba[6].ToLAB(), v_rgba[7].ToLAB()},
		},
	}
	const steps = 6
	var output_OkLAB [][][]cs.OkLAB
	// output_OkLAB = _trilerp(verts, steps*90)
	// imgs := export.Export_Cube(output_OkLAB)
	// for idx, img := range imgs {
	// 	export.Save_PNG(img, "./images/"+strconv.Itoa(idx)+".png")
	// }

	output_OkLAB = _trilerp(verts, steps)
	output := export.Ansi_Cube(output_OkLAB, export.AnsiOpts{Format: "%d%d%d", Indices: true, Spacing: 3})
	fmt.Print(output)
	// fmt.Println(time.Since(start))
	// fmt.Printf("%#v\n", os.Args)
}

// vim: set noet ts=4 sw=4 ff=unix
