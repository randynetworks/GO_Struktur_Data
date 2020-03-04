package main

import "fmt"

// array

func main() {

	member := [5]string{"Randy", "Paul", "naruto", "sasuke", "rikudou"}
	sliceMember := member[:3]

	fmt.Println(sliceMember)

	// ADA LENGHT DARI ARRAY
	fmt.Println("Lengthnya", len(sliceMember))
	// ADA KAPASITAS ARRAY, NAMUN IA MENGABIL UKURAN KAPASITAS DATA UTAMANYA
	fmt.Println("Kapasitasnya", cap(sliceMember))

	// ADA FUNGSI BERNAMA make(), berguna untuk membooking tempat array kosong
	// parameter yang diisi make( []tipe data arraynya, length, capacity )
	anggota := make([]string, 5, 10)
	fmt.Println(anggota)
	fmt.Println("Lengthnya", len(anggota))
	fmt.Println("Kapasitasnya", cap(anggota))
}
