package main

import "fmt"

func main() {

	var berat, totalBiaya, kirimKg, kirimGr int

	fmt.Print("Berat parsel (gram) : ")
	fmt.Scanln(&berat)

	beratKg := berat / 1000
	sisaGr := berat % 1000
	kirimKg = beratKg * 10000

	if beratKg <= 10 {
		if sisaGr >= 500 {
			kirimGr = sisaGr * 5
		} else {
			kirimGr = sisaGr * 15
		}
	}

	totalBiaya = kirimKg + kirimGr

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGr)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", kirimKg, kirimGr)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
