package main

import "fmt"

func main() {
	var berat, hasil, ongkirkg, ongkirgram int
	fmt.Scan(&berat)

	kg := berat / 1000
	gram := berat % 1000

	if berat < 10000 {
		ongkirkg = kg * 10000
		if gram > 500 {
			ongkirgram = gram * 5
			hasil = ongkirkg + ongkirgram
		} else {
			ongkirgram = gram * 15
			hasil = ongkirkg + ongkirgram
		}
	} else {
		ongkirkg = kg * 10000
		if gram > 500 {
			ongkirgram = gram * 5
			hasil = ongkirkg + ongkirgram
		} else {
			ongkirgram = gram * 15
			hasil = ongkirkg + ongkirgram
		}
		hasil = ongkirkg
	}
	fmt.Println("Berat Parsel: ", berat)
	fmt.Println("Detail Berat: ", kg, "kg + ", gram, "gr ")
	fmt.Println("Detail biaya:", "Rp. ", ongkirkg, " + Rp.", ongkirgram)
	fmt.Print(hasil)
}
