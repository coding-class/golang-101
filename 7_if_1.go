package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 10
	c := 20

	if a + b > c {
		fmt.Println(`1 -> a + b lebih besar daripada c`)
	} else if a + b == c {
		fmt.Println("2 -> a + b sama dengan c")
	} else {
		fmt.Println("3 -> a + b lebih kecil daripada c")
	}

	if c % 2 == 0 {
		fmt.Println("c adalah bilangan genap")
	} else {
		fmt.Println("c adalah bilangan ganjil")
	}

	// if a + b == c {
	// 	fmt.Println("Benar")
	// } else {
	// 	fmt.Println("Salah")
	// }

	// if ...(kondisi).. {
	// 	(aksi yang dilakukan, jika kondisi terpenuhi / true)
	// }
}
