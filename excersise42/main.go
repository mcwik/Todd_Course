package main

import "fmt"

func main() {

	ar1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range ar1 {

		fmt.Printf("The value: %v is type of: %T\n", v, v)

	}

	fmt.Printf("The slice variable is type of: %T\n", ar1)

	fmt.Printf("Length of slice is: %v\n", len(ar1))
	fmt.Printf("Capacity of slice is: %v\n", cap(ar1))

	fmt.Printf("%#v\n", ar1)

}
