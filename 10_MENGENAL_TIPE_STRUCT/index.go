package main

import "fmt"

// STRUCT
type member struct {
	name, job string
	age       int
}

func main() {
	member1 := member{
		name: "Randy",
		job:  "Programmer",
		age:  10,
	}
	// value type = jika di assign tidak akan menganggu data asli

	member2 := member1
	member2.name = "Randy Ramadhan"

	fmt.Println(member1)
	fmt.Println(member2)
	fmt.Println("Nama Orang ini " + member1.name)

}
