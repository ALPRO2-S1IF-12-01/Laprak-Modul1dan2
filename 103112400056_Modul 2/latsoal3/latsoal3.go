// Felix Pedrosa V

package main

import "fmt"

func main() {

	var berat, beratKg, sisaGram, biayaKg, biayaSisaGram, totalBiaya int

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&berat)

	beratKg = berat / 1000
	sisaGram = berat % 1000

	biayaKg = beratKg * 10000

	if sisaGram > 0 {
		if beratKg > 10 {
			biayaSisaGram = 0
		} else if sisaGram >= 500 {
			biayaSisaGram = sisaGram * 5
		} else {
			biayaSisaGram = sisaGram * 15
		}
	}

	totalBiaya = biayaKg + biayaSisaGram

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisaGram)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
