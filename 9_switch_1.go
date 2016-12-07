package main

import (
	"fmt"
)

func main() {
	nilaiHuruf := "A"

	switch nilaiHuruf {
	case "A":
		fmt.Println("Sangat Bagus")
	case "B":
		fmt.Println("Bagus")
	case "C":
		fmt.Println("Lumayan")
	case "D":
		fmt.Println("Tingkatkan, Bro!")
	default:
		fmt.Println("Hmmm")
	}
}
