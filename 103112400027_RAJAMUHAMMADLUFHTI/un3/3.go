// RAJA MUHAMMAD LUFHTI || 103112400027 
package main

import "fmt"

func main() {
	var berat, hasil, hargaperkg, hargapergram int
	fmt.Scan(&berat)

	kg := berat / 1000
	gram := berat % 1000

	if berat < 10000 {
		hargaperkg = kg * 10000
		if gram > 500 {
			hargapergram = gram * 5
			hasil = hargaperkg + hargapergram
		} else {
			hargapergram = gram * 15
			hasil = hargaperkg + hargapergram
		}
	} else {
		hargaperkg = kg * 10000
		if gram > 500 {
			hargapergram = gram * 5
			hasil = hargaperkg + hargapergram
		} else {
			hargapergram = gram * 15
			hasil = hargaperkg + hargapergram
		}
		hasil = hargaperkg
	}
	fmt.Println("Berat Parsel: ", berat)
	fmt.Println("Detail Berat: ", kg, "kg + ", gram, "gr ")
	fmt.Println("Detail biaya:", "Rp. ", hargaperkg, " + Rp.", hargapergram)
	fmt.Print(hasil)
}
