package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var x int

	for i := 0; i < 42; i++ {

		x = rand.Intn(5)

		fmt.Println("iterator: ", i)
		fmt.Print(x, "- ")

		switch x {
		case 0:
			fmt.Println("To jest zero")
		case 1:
			fmt.Println("To jest jeden")
		case 2:
			fmt.Println("To jest dwa")
		case 3:
			fmt.Println("To jest trzy")
		case 4:
			fmt.Println("To jest cztery")
		}
	}

}
