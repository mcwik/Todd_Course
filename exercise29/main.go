package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Println(x)
	fmt.Println(y)

	for i := 0; i < 100; i++ {

		fmt.Println(i)
	}

	for i := 0; i < 100; i++ {

		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("X: %v\n", x)
		fmt.Printf("Y: %v\n\t", y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both are less than four")
		case x > 6 && y > 6:
			fmt.Println("Both are grater than six")
		case x >= 4 && x <= 6:
			fmt.Println("x is between 4 and 6. Both included.")
		case y != 5:
			fmt.Println("y is not five")
		default:
			fmt.Println("None of the above")
		}
	}
	fmt.Println(x)
	fmt.Println(y)

}
