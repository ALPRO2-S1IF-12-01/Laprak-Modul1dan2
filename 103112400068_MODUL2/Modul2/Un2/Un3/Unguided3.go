package main

import "fmt"

func main() {
	var Berat int
	fmt.Print("Berat Parsel (gram): ")
	fmt.Scan(&Berat)

	kg := Berat / 1000
	gr := Berat % 1000

	biayaKg := kg * 10000
	var biayaGr int

	if kg > 10 {
		biayaGr = 0
	} else if gr >= 500 {
		biayaGr = gr * 5
	} else {
		biayaGr = gr * 15
	}

	totalBiaya := biayaKg + biayaGr

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGr)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
