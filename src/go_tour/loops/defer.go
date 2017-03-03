package main

import "fmt"

func main() {
	a := 1
	b := 2
	defer fmt.Println(a)
	a = a + b
	fmt.Println(a)
	defer fmt.Println(a)
	defer fmt.Println("world")
	x:= 4
	defer fmt.Println(x)
	fmt.Println("hello")
}