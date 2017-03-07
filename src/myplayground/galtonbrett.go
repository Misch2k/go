package main

import (
	"fmt"
	"math/rand"
)

func main(){

	var temp, input int = 0,0

	fmt.Print("Wie hoch soll das Galtonbrett sein? \n")
	fmt.Scanln(&input)
	height := input

	fmt.Print("Wie viele Bälle sollen durch laufen? \n")
	fmt.Scanln(&input)
	balls := input

	can := make([] int, height+1)

	for i:= 0; i<height+1; i++ {
		fmt.Print("   ")
		for k:=0; k<2*height-2*i; k++{
			fmt.Print(" ")
		}
		for j:= 0; j<i; j++ {
			fmt.Print(".   ")
		}
		fmt.Println( "")
	}

	for i:=0; i<balls; i++ {
		for j:=0; j<height; j++{
		temp += rand.Intn(2)
		}
		can[temp]++
		temp = 0
	}

	for i:=0; i<height+1; i++ {
		fmt.Print("|")
		if can[i] < 10 {
			fmt.Print(can[i], "| ")
		}
		if can[i] >= 10 && can[i] < 100 {
			fmt.Print(can[i], "|")
		}
		if can[i] >= 100 && can[i] < 1000 {
			fmt.Print(can[i], "|")
		}

	}
	fmt.Println("")
	for i:=0; i<height+1; i++ {
		fmt.Print(" ")
		if can[i] < 10 {
			fmt.Print("\u035E", "  ")
		}
		if can[i] >= 10 && can[i] < 100 {
			fmt.Print("\u0350E\u0350E", " ")
		}
		if can[i] >= 100 && can[i] < 1000 {
			fmt.Print("---", "")
		}

	}

}