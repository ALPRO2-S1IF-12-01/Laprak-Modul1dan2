package main

import (
	"fmt"
	"strings"
)

func tampilkanGaris() {
	fmt.Println(strings.Repeat("=", 50))
}
func hitungBiayaParsel(beratGram int) (int, int, int, int, int) {
	kg := beratGram / 1000
	gram := beratGram % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	if gram >= 500 {
		biayaSisa = gram * 5
	} else if kg <= 10 && gram > 0 {
		biayaSisa = gram * 15
	}

	totalBiaya := biayaKg + biayaSisa
	return kg, gram, biayaKg, biayaSisa, totalBiaya
}

func main() {
	var beratGram int
	tampilkanGaris()
	fmt.Println("PROGRAM PERHITUNGAN BIAYA PARSEL")
	tampilkanGaris()
	fmt.Print("\nMasukkan berat parsel (gram): ")
	fmt.Scan(&beratGram)
	if beratGram <= 0 {
		fmt.Println("\nError: Berat parsel harus lebih besar dari 0!")
		return
	}
	kg, gram, biayaKg, biayaSisa, totalBiaya := hitungBiayaParsel(beratGram)
	tampilkanGaris()
	fmt.Println("DETAIL PERHITUNGAN")
	tampilkanGaris()
	fmt.Printf("Berat parsel\t: %d gram\n", beratGram)
	fmt.Printf("Konversi\t: %d kg + %d gram\n", kg, gram)
	fmt.Println("\nRincian biaya:")
	fmt.Printf("- Biaya per kg\t: Rp. %d\n", biayaKg)
	fmt.Printf("- Biaya sisa\t: Rp. %d\n", biayaSisa)
	tampilkanGaris()
	fmt.Printf("Total biaya\t: Rp. %d\n", totalBiaya)
	tampilkanGaris()
	fmt.Println("\nCatatan:")
	fmt.Println("- Biaya per kg: Rp. 10.000")
	fmt.Println("- Biaya sisa ≥ 500 gram: Rp. 5/gram")
	fmt.Println("- Biaya sisa < 500 gram: Rp. 15/gram (jika total ≤ 10 kg)")
	tampilkanGaris()
}
