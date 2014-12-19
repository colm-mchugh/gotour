package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

// Calculate square root using delta
// Return error if given a negative number
func Sqrt(x float64) (float64, error) {
	z, i := x, 0
	const delta = .0000001
	if z < 0 {
		return x, ErrNegativeSqrt(x)
	}
	for {
		i++
		z_prev := z
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-z_prev) < delta {
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
