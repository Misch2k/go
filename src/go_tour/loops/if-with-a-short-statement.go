package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else { //else must bee on the same line then the } tag!
	return v
	}
	//return v <- Not allowed : v is only in IF loop and else loop valid
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}