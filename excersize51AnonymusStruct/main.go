package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friend    map[string]int
		favDrinks []string
	}{
		first:     "Jeremy",
		friend:    map[string]int{"Ritchard": 1, "James ": 2},
		favDrinks: []string{"Bear", "Gin", "Rum"},
	}

	fmt.Println(p1.first)

	fmt.Println("His friends:")
	fmt.Println(p1.friend)

	fmt.Println("His fav drinks:")
	fmt.Println(p1.favDrinks[0])
	fmt.Println(p1.favDrinks[1])
	fmt.Println(p1.favDrinks[2])

}
