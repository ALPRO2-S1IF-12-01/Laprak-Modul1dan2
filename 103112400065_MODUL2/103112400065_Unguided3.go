package main

import "fmt"

func main() {
	var berat, kg, gr, harga_kg, harga_gr, totalbiaya int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)
	kg = berat / 1000
	gr = berat % 1000

	if gr >= 500 {
		harga_gr = gr * 5
	} else {
		harga_gr = gr * 15
	}

	if kg <= 10 {
		harga_kg = kg * 10000
		totalbiaya = harga_kg + harga_gr
	} else {
		harga_kg = kg * 10000
		totalbiaya = harga_kg
	}

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
	fmt.Printf("Detail biaya: Rp. %d  + Rp. %d\n", harga_kg, harga_gr)
	fmt.Printf("Total biaya: Rp. %d", totalbiaya)
}
