// MULIA AKBAR NANDA PRATAMA
// 103112400034
package main

import "fmt"

func main() {
	var berat_gram int

	fmt.Print("masukkan berat parsel (gram): ")
	fmt.Scan(&berat_gram)

	kg := berat_gram / 1000
	Gram := berat_gram % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	if Gram >= 500 {
		biayaSisa = Gram * 5
	} else {
		if kg <= 10 && Gram > 0 {
			biayaSisa = Gram * 15
		}
	}

	totalBiaya := biayaKg + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, Gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
