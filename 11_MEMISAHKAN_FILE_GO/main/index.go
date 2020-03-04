package main

import (
	"fmt"
	// cara mengimport file member.go
	"github.com/randynetworks/belajar_go/stuktur_data_go/11_MEMISAHKAN_FILE_GO/models"
)

// untuk memisahkan file dan menghubungkannya, di haruskan, setiap nama data struct dan value nya harus diawali dengan huruf kapital

func main() {
	// cara aksesnya tinggal tambahkan nama package(disesuaikan)nya, diikuti dengan nama struct nya
	member1 := models.Member{
		Name: "Randy",
		Job:  "Programmer",
		Age:  10,
	}
	fmt.Println(member1)

}
