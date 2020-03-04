package main

import "fmt"

// array

func main() {
	// array punya saudara yaitu slices

	// slice
	// perbedaan mencokok, array dicantumkan dengan array, sedangkan slices engga
	member := [5]string{"Randy", "Paul", "naruto", "sasuke", "rikudou"}

	// membuat slices dari array yang sudah ada
	// jika hanya [:] maka seluruh data diambil
	// sliceMember := member[:]

	// jika ingin spesifik, cantumkan index nya diantara [indexAwal:IndexAkhir]
	// tapi yang akan di ambil indexAwal, masuk, sedangkan indexAkhir tidak akan di masukan (mengabil index sebelumnya)
	// begitu pula salahsatunya
	sliceMember := member[:4]

	// perlu diperhatikan,, slices dapat menimpa seluruh data originalnya

	fmt.Println(member)
	fmt.Println(sliceMember)

}
