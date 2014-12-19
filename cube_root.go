package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := 1.0
	x_r := real(x)
	for i := 0; i < 1000; i++ {
		z = z - ((z*z*z)-x_r)/(3*(z*z))
	}
	return complex(z, imag(x))
}

func main() {
	fmt.Println(Cbrt(2 + 3i))
	fmt.Println(cmplx.Pow(2, 0.333333))
}
