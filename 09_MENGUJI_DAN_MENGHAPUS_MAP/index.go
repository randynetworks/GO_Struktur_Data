package main

import "fmt"

var member = map[int]string{
	10: "randy",
	20: "Naruto",
	30: "Sasuke",
}

func main() {
	// map
	// initialisasi kosong dengan map
	// name := map[int]string{}
	// name := make(map[int]string)

	// UNTUK MENGHAPUS DATA
	// parameter yang disini delete( namaMapnya, key)
	delete(member, 30)
	fmt.Println(member)
	if checkMember(40) {
		fmt.Println("Data Ada")
	} else {
		fmt.Println("Data tidak ada")
	}
}

func checkMember(id int) bool {
	_, isExsist := member[id]
	return isExsist
}
