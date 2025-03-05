// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func main() {
	var berat, kg, gr, biayaKg, biayaGr, total int
	fmt.Scan(&berat)
	kg = berat / 1000
	gr = berat % 1000
	biayaKg = kg * 10000

	if gr >= 500 {
		biayaGr = gr * 5
	} else {
		biayaGr = gr * 15
	}
	
	if kg > 10 {
		total = biayaKg
	} else {
		total = biayaKg + biayaGr
	}
	fmt.Printf("Berat berat (gram): %v\n", berat)
	fmt.Printf("Detail berat: %v kg + %v gr\n", kg, gr)
	fmt.Printf("Detail biaya: Rp. %v + Rp. %v\n", biayaKg, biayaGr)
	fmt.Printf("Total biaya: Rp. %v", total)
}