package main

import "fmt"

func main() {
	// map = teknik data key : value
	member := map[int]string{
		10: "randy",
		20: "Naruto",
		30: "Sasuke",
	}

	// reference type = nilai original akan berubah
	newMember := member
	// merubah data
	newMember[10] = "Rikudou"

	fmt.Println(newMember)
	fmt.Println(newMember[10])

	// untuk mengeluarkan datanya 1 persatu, harus menggunakan range

	for id, name := range member {
		fmt.Println(id, name)
	}
}
