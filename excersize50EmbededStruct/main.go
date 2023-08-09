package main

import (
	"fmt"
)

func main() {

	type engine struct {
		electric bool
	}

	type vechicle struct {
		engine
		make  string
		model string
		dors  int
		color string
	}

	vech1 := vechicle{

		engine: engine{electric: true},
		make:   "Tesla",
		model:  "Roadster",
		dors:   3,
		color:  "Yellow",
	}

	vech2 := vechicle{

		engine: engine{electric: false},
		make:   "Porshe",
		model:  "Boxter",
		dors:   3,
		color:  "Silver",
	}

	vechicles := []vechicle{vech1, vech2}

	for _, v := range vechicles {

		fmt.Println("Company:", v.make)
		fmt.Println("Model:", v.model)
		fmt.Println("Dors:", v.dors)
		fmt.Println("Color", v.color)
		fmt.Println("Is electric:", v.electric)
		fmt.Println("*********************************")

	}

}
