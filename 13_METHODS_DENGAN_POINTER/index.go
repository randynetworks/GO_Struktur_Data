package main

import "fmt"

// METHODS
// MIRIP DNEGAN FUNGSI, TAPI DI SPECIAL/ MILIK DARI OBJECT TERTENTU

type member struct {
	name string
	age  int
}

func (m *member) getInfo(newName string, newAge int) {
	m.name = newName
	m.age = newAge
}

func main() {
	user1 := member{
		name: "randy",
		age:  90,
	}

	user1.getInfo("Paul", 20)
	fmt.Println(user1)
}
