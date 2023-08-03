package main

import (
	"fmt"
)

func main() {

	sl1 := []int{1, 2, 3, 4, 5, 6, 7}
	sl2 := sl1

	fmt.Println(sl1)
	fmt.Printf("%p", sl1)
	fmt.Println(sl2)
	fmt.Printf("%p", sl2)

}
