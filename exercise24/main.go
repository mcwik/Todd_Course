package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(250)

	fmt.Printf("Variable x has value %v\n", x)

	if x <= 100 {
		fmt.Println("x between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("x between 101 and 200")
	} else {
		fmt.Println("x between 201 and 249")
	}

	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

}
