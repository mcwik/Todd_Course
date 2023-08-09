package main

import "fmt"

func main() {

	jbm := [][]string{{"James", "Bond", "Shaken, not stired"}, {"Miss", "Moneypenny", "I'm 008"}}

	/*
		"James", "Bond", "Shaken, not stired"
		"Miss", "Moneypenny", "I'm 008"
	*/

	for i, vex := range jbm {

		fmt.Println("**************", i, "******************")

		for j, vin := range vex {

			switch j {
			case 0:
				fmt.Println("Name:", vin)
			case 1:
				fmt.Println("Surname:", vin)
			case 2:
				fmt.Println("Motto:", vin)
			}

		}

		fmt.Println("***********************************")

	}

}
