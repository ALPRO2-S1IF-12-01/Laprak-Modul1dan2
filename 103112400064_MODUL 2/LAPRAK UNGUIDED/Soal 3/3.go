//MUHAMMAD FAUZAN
package main

import (
	"fmt"
)

func main() {
	var berat, hargaGram int
	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scanln(&berat)
	kilogram := berat / 1000
	sisaGram := berat % 1000
	fmt.Printf("Detail Berat: %d kg dan %d gr\n", kilogram, sisaGram)
	hargaKg := kilogram * 10000

	switch {
	case sisaGram >= 500:
		hargaGram = sisaGram * 5
	default:
		hargaGram = sisaGram * 15
	}

	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", hargaKg, hargaGram)
	fmt.Printf("Total Biaya: Rp. %d\n", hargaKg+hargaGram)
}
