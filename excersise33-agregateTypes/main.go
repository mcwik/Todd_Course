package main

import (
	"fmt"
)

func main() {

	x := []string{"Marcel", "Krzywy", "Brzydal", "Kutek", "Maniek"}

	fmt.Printf("address of the slice: %p\n", x)

	x = append(x, "Smorawa", "Marian Wielun")

	fmt.Printf("address of the slice: %p\n", x)

	for _, v := range x {

		fmt.Println(v)
	}

	fmt.Println([]string{"Marcel", "Krzywy", "Brzydal", "Kutek", "Maniek"})

}
