package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for delta := 1.0; delta > .001; {
		oldZ := z
		z = z - ((z*z)-x)/2*z
		delta = math.Abs(oldZ - z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
