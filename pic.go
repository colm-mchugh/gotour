package main

import "code.google.com/p/go-tour/pic"

func foo(x, y int) int {
	return x * y
}

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		r[i] = make([]uint8, dx)
	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			r[y][x] = uint8(foo(x, y))
		}
	}

	return r
}

func main() {
	pic.Show(Pic)
}
