//Ahmad Ruba'i
//103112400074

package main

import "fmt"

func main() {
	var beratGram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratGram)

	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaPerKg := 10000
	biayaKg := beratKg * biayaPerKg
	biayaGram := 0

	if beratKg <= 10 {
		if sisaGram < 500 {
			biayaGram = sisaGram * 15
		} else {
			biayaGram = sisaGram * 5
		}
	}

	totalBiaya := biayaKg + biayaGram

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGram)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
