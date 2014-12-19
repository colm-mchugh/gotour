package main

import (
	"fmt"
	"math"
)

// Calculate square root using delta
func Sqrt1(x float64) (z float64, i int64) {
	z, i = x, 0
	const delta = .0000001
	for {
		i++
		z_prev := z
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-z_prev) < delta {
			return
		}
	}
}

// Calculate square root using fixed # iterations
func Sqrt2(x float64) (float64, int8) {
	z := x
	const max_i = 10
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z, max_i
}

func main() {
	fmt.Println(Sqrt1(3))
	fmt.Println(Sqrt2(3))
	fmt.Println(math.Sqrt(3))
}
