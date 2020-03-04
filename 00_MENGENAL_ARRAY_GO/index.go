package main

import "fmt"

// array

func main() {

	var member [3]string
	member[0] = "Randy"
	member[1] = "Naruto"
	member[2] = "Paul"

	// angka := [4]int{1, 4, 3, 2}
	//atau agar lebih fleksibel dengan ukuran index
	angka := [...]int{1, 4, 3, 2, 4, 3, 2}

	angka[1] = 100

	fmt.Println(angka)
}
