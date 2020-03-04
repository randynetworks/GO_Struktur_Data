package main

import "fmt"

// array

func main() {
	// array Multidimensi

	hewan := [3][2]string{
		{"Macan", "singa"},
		{"hiu", "paus"},
		{"elang", "gagak"},
	}
	// [00] [01]
	// [10] [11]
	// [20] [21]
	fmt.Println(hewan[2][1])

}
