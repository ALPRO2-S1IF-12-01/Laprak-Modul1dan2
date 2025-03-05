package main

import "fmt"

func main() {
	var beratParsel, kg, gr, biayaKg, biayaGr, totalBiaya int

    fmt.Scan(&beratParsel)
    kg = beratParsel / 1000
    gr = beratParsel % 1000

    if gr >= 500 {
        biayaGr = gr * 5
    } else {
        biayaGr = gr * 15
    }

    biayaKg = kg * 10000
    if kg > 10 {
		totalBiaya = biayaKg
	} else {
		totalBiaya = biayaKg + biayaGr
	}

	fmt.Printf("Berat parsel (gram): %v \nDetail berat: %v kg + %v gr \nDetail biaya: Rp. %v + Rp. %v \nTotal biaya: Rp. %v", beratParsel, kg, gr, biayaKg, biayaGr, totalBiaya)
}