package main

import (
	"fmt"
)

func main() {
	// pengulangan menaik
	for i := 1; i <= 50; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	// pengulangan menurun
	for i := 50; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
	// pengulangan tak terhingga (infinity loop)
	// for i := 1; i <= 1; i-- {
	// 	fmt.Print(i, " ")
	// }
}
