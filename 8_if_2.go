package main

import (
	"fmt"
)

func main() {
	nilai := "A"

	if nilai == "A" {
		fmt.Println("Sangat Bagus")
	} else if nilai == "B" {
		fmt.Println("Bagus")
	} else if nilai == "C" {
		fmt.Println("Lumayan")
	} else if nilai == "D" {
		fmt.Println("Tingkatkan, Bro!")
	} else {
		fmt.Println("Hmmm")
	}
}
