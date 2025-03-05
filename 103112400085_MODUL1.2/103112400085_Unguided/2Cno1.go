package main

import (
	"fmt"
)

func hitungBiaya(beratGram int) (int, int, int, int, int) {
	kg := beratGram / 1000
	sisaGram := beratGram % 1000

	var biayaKg, biayaSisa int

	// Hitung biaya per kg
	if kg > 0 {
		biayaKg = kg * 10000
	} else {
		biayaKg = 0
	}

	// Hitung biaya tambahan untuk sisa gram
	if sisaGram > 0 {
		biayaSisa = (sisaGram / 100) * 500
	} else {
		biayaSisa = 0
	}

	// Total biaya
	totalBiaya := biayaKg + biayaSisa
	return kg, sisaGram, biayaKg, biayaSisa, totalBiaya
}

func main() {
	var berat int
	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&berat)

	if berat < 0 {
		fmt.Println("Berat tidak boleh negatif!")
	} else {
		kg, sisaGram, biayaKg, biayaSisa, totalBiaya := hitungBiaya(berat)

		fmt.Printf("Berat parsel (gram): %d\n", berat)
		fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
		fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
	}
}