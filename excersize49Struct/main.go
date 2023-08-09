package main

import "fmt"

func main() {

	type person struct {
		fname  string
		sname  string
		age    int
		skills []string
	}

	p1 := person{
		fname:  "Benjamin",
		sname:  "Franklin",
		age:    42,
		skills: []string{"Writing poems", "Electricity"},
	}

	p2 := person{
		fname:  "Martha",
		sname:  "Graham",
		age:    62,
		skills: []string{"Creativity", "Modern dance choreographer"},
	}

	//persons := []person{p1, p2}

	mpersons := map[string]person{p1.sname: p1, p2.sname: p2}

	for _, v := range mpersons {

		fmt.Println(v.fname)
		fmt.Println(v.sname)
		fmt.Println(v.age)
		for j, s := range v.skills {

			fmt.Printf("\t %v. %v\n", j+1, s)

		}
	}

	/*
		for _, v := range persons {

			fmt.Println(v.fname)
			fmt.Println(v.sname)
			fmt.Println(v.age)
			fmt.Println("Skills:")

			for j, s := range v.skills {

				fmt.Printf("\t %v. %v\n", j+1, s)

			}

		}
	*/

}
