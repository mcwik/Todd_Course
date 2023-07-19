package main

import (
	"fmt"
)

func main() {

	//x := []int{1, 2, 3, 4, 5}

	m := map[string]int{
		"james":      42,
		"moneypenny": 32,
	}

	fmt.Printf("Checking the value of key 'Q': %v\n\n", m["Q"])

	for k, v := range m {

		//fmt.Printf("The value is: %v and the index is: %v\n", v, i)

		fmt.Println(k)
		fmt.Println(v)

		/*
			"comma ok" idiom;
			Below code chcecks if the 'm' map stores the value with key 'Q'
			1. ok := m["Q"] stores boolean value -> true if 'Q' key exists in the map and fales if it doesn't
			2. If exists -> it save the value in the 'v' variable
			3. ok will store ture or false
		*/

		if v, ok := m["james"]; ok {
			fmt.Printf("Q key exist int the map, it's value is %v\n", v)
		} else {
			fmt.Printf("Q key does not exist in the map %v\n\n", v)
		}

		/*
			if v, ok := m["Q"]; !ok { //DOES NOT OK!
				fmt.Printf("Q key does not exist in the map %v\n\n", v)
			} else {
				fmt.Printf("Q key exist int the map, it's value is %v", v)
			}
		*/

	}

}
