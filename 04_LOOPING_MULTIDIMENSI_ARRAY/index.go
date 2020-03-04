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
	for i := 0; i < len(hewan); i++ {
		for j := 0; j < len(hewan[i]); j++ {
			fmt.Println(hewan[i][j])
		}
		fmt.Println("============")
	}

}
