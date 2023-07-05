package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func main() {

	data, err := os.ReadFile("SNOWY-EVENING.txt")
	if err != nil {
		log.Fatal(err)
	}
	//os.Stdout.Write(data)

	sum := sha256.Sum256(data)

	fmt.Printf("%x", sum)

}
