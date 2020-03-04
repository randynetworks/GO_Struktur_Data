package main

import "fmt"

// array

func main() {

	member := []string{"Randy", "Paul", "naruto", "sasuke", "rikudou"}
	// sliceMember := member[:3]

	newMember := make([]string, 5, 10)
	// FUNGSI copy() berguna agar ketika kita menggambil data utamanya, dan jika diubah, data utamanya tidak akan tertimpa
	// parameter yang diisi adalah copy( kemanaDataYangAkanDiTimpa, DataUtamanya)
	copy(newMember, member)
	newMember[1] = "george"

	// jika kita akan menambahkan data baru, gunakan fungsi append()
	// parameter yang diisi adalah append(arraynya, nilainya, .. )
	newMember = append(newMember, "Ringgo", "John")

	// TAHAPAN JIKA INGIN MEMANIPULASI DATA ARRAY/SLICES
	// 1. Membuat array kosong baru dengan fungsi make()
	// 2. mengkopi data ke dalam array baru dari array utama (agar tidak menganggu data array utama) dengan copy()
	// 3. menambahkan data dengan append()

	fmt.Println(newMember)
	fmt.Println(member)
}
