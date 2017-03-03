//With the Newton-rule: z = z-((z^2-x)/2z)
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64, lim int) float64 {
	z := x
	for i := 0; i < lim ; i++{
		z = (math.Pow(z,2) + x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2, 5))
	fmt.Println(math.Sqrt(2))
}
