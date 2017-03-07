package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  	// has type Vertex
	v2 = Vertex{X: 1}  		// Y:0 is implicit
	v3 = Vertex{}      		// X:0 and Y:0
	p  = &Vertex{1, 2} 	// has type *Vertex
)

func main() {
	p.X = 5
	fmt.Printf("p is type of %T\n", p)
	fmt.Printf("v1.X is type of %T\n", v1.X)
	fmt.Printf("v2 is type of %T\n", v2)
	fmt.Printf("v3 is type of %T\n", v3)
	fmt.Println(v1, p, v2, v3)
}
