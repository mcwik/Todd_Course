package main

import (
	"fmt"
	"math/rand"
)

func init() {

	fmt.Println("This is where initialization for my program occurs")
}

func init() {

	//fmt.Println("This is where initialization for my program occurs2")
}

func main() {

	x := rand.Intn(250)

	fmt.Printf("Variable x has value %v\n", x)

	switch {
	case x <= 100:
		fmt.Println("x between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("x between 101 and 200")
	case x >= 201 && x < 250:
		fmt.Println("x between 201 and 249")
	}

}
