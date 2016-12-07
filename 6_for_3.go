package main

import "fmt"

func main() {
	go printEA()
	go printOOOO()
	go printWEW()

	var stop string
	fmt.Scanf("%s", &stop)
}

func printEA() {
	for {
		fmt.Print("EA")
	}
}

func printWEW() {
	for {
		fmt.Print("WEW")
	}
}

func printOOOO() {
	for {
		fmt.Print("0000")
	}
}
