// PRATAMA BINTANG DANISWARA
// 103112400051
package main

import "fmt"

func main() {
	var  beratgr, sisa int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratgr)

	parsel := beratgr / 1000
	berat := beratgr % 1000
	fmt.Printf("Detail berat: %d kg + %d gr\n", parsel, berat)

	biaya := parsel * 10000
	if berat < 500 {
		sisa = berat * 15
	} else if berat >= 500 {
		sisa = berat * 5
	}

	fmt.Printf("Detail biaya: Rp. %d + Rp. %d \n", biaya, sisa)
	fmt.Printf("Total biaya: Rp. %d", biaya+sisa)
}
