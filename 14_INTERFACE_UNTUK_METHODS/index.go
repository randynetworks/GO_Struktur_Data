package main

import (
	"fmt"
	"math"
)

// INTERFACE -> Kumpulan dari methods yang kita buat

type bentuk interface {
	keliling() float64
	luas() float64
}

type kotak struct {
	panjang, lebar float64
}

func (k kotak) keliling() float64 {
	return 2*k.panjang + 2*k.lebar
}
func (k kotak) luas() float64 {
	return k.panjang * k.lebar
}

type lingkaran struct {
	radius float64
}

func (l lingkaran) keliling() float64 {
	return math.Pi * 2 * l.radius
}
func (l lingkaran) luas() float64 {
	return math.Pi * l.radius * l.radius
}
func main() {
	kotak1 := kotak{5, 10}
	hitungBangungan(kotak1)

	lingkaran1 := lingkaran{7}
	hitungBangungan(lingkaran1)
}

func hitungBangungan(b bentuk) {
	fmt.Println("Kelilingnya:", b.keliling())
	fmt.Println("luasnya:", b.luas())
}
