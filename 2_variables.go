package main

import (
	"fmt"
)

func main() {
	// DEKLARASI variabel dengan tipe data bilangan (integer)
	// var a int // default value = 0
	// var b = int(1)
	// c := int(2)
	// fmt.Println(a, b, c)

	// a = b + c
	// fmt.Println(a, b, c)

	// STRING (KATA)
	// var kata1 string
	// var kata2 = "ini kata 2"
	// kata3 := "ini kata 3"
	// fmt.Println(kata1)
	// fmt.Println(kata2)
	// fmt.Println(kata3)

	// fmt.Println("======")

	// kata1 = kata2 + kata3
	// fmt.Println(kata1)
	// fmt.Println(kata2)
	// fmt.Println(kata3)

	float1 := float32(3.14159)	// bilangan desimal
	float2 := float64(22.0/7.0)	// bilangan desimal
	float3 := float64(22/7)	// bilangan desimal
	bool1 := bool(true)			// tipe data BOOLEAN -> benar atau salah
	bool2 := bool(false)		// tipe data BOOLEAN -> benar atau salah
	var a int // default value = 0
	var b = int(1)

 	a = float1 + b
 	float2 = b

 	fmt.Println(a, float2)
	fmt.Println(float1, float2, float3)
	fmt.Println(bool1, bool2)

	// bool3 := bool1 && bool2
	// fmt.Println("OPERASI BOOLEAN AND")
	// fmt.Println(true && true)
	// fmt.Println(true && false)
	// fmt.Println(false && true)
	// fmt.Println(false && false)

	// fmt.Println("OPERASI BOOLEAN OR")
	// fmt.Println(true || true)
	// fmt.Println(true || false)
	// fmt.Println(false || true)
	// fmt.Println(false || false)
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
