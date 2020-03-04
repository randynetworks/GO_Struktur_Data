package main

import "fmt"

// METHODS
// MIRIP DNEGAN FUNGSI, TAPI DI SPECIAL/ MILIK DARI OBJECT TERTENTU

type member struct {
	name string
	age  int
}

func (m member) getInfo() {
	fmt.Println("ini diprint member")
}

func main() {
	user1 := member{
		name: "randy",
		age:  90,
	}
	user1.getInfo()
}
