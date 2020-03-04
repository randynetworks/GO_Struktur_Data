package main

import "fmt"

// array

func main() {

	member := [3]string{"Randy", "Paul", "Naruto"}

	// PRINT MANUAL
	// fmt.Println(member[0])
	// fmt.Println(member[1])
	// fmt.Println(member[2])

	// dengan for
	for i := 0; i < len(member); i++ {
		fmt.Println(member[i])
	}
}
