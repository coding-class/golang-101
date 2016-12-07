package main

import (
	"fmt"
)

func main() {
	fmt.Println("Halo, nama kamu siapa?")

	var nama string
	fmt.Scanf("%s", &nama)

	fmt.Println("Halo," , nama)
}
