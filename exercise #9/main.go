package main

import "fmt"

var x int = 20

const y int = 10

func main() {

	z := 30

	fmt.Println("Package scope vairable x: ", x)
	fmt.Println("Package scope constant y: ", y)

	fmt.Println("Block scope variable z: ", z)

}
