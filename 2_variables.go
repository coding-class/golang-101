package main

import (
	"fmt"
)

func main() {
	// variabel dengan tipe data bilangan (integer)
	var a int
	var b = int(1)
	c := int(2)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var kata1 string
	var kata2 = "ini kata 2"
	kata3 := "ini kata 3"

	fmt.Println(kata1)
	fmt.Println(kata2)
	fmt.Println(kata3)
}

// cara deklarasi variabel di golang
/*
var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map lookup; only interested in "found"
*/
