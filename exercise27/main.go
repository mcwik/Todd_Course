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

	if x < 4 && y < 4 {
		fmt.Println("Both are less than four")
	} else if x > 6 && y > 6 {
		fmt.Println("Both are grater than six")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is between 4 and 6. Both included.")
	} else if y != 5 {
		fmt.Println("y is not five")
	} else {
		fmt.Println("None of the above")
	}
}
