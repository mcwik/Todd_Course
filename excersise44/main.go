package main

import "fmt"

func main() {

	ar := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	ar1 := ar[:5]
	ar2 := ar[5:]
	ar3 := ar[2:7]
	ar4 := ar[1:6]

	sl := [][]int{ar1, ar2, ar3, ar4}

	fmt.Println(ar1)
	fmt.Println(ar2)
	fmt.Println(ar3)
	fmt.Println(ar4)

	fmt.Println("*******************************************'")

	for _, vex := range sl {

		fmt.Println(vex)

	}

}
