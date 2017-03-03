package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	//oder
	for ; sum < 1000 ; {
		sum++
	}
	fmt.Println(sum)

	//oder
	for sum < 10000 {
		sum++
	}

	/* Endlosschleife
	for {
		sum++
		fmt.Println(sum)
	}
	*/
	fmt.Println(sum)
}
