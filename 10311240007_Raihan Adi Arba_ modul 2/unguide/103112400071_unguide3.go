//Nama	: Raihan Adi Arba
//NIM	: 103112400071
//KELAS	: IF-12-011

package main

import "fmt"

func main() {
	var berat_parsel, detail_biaya, total_berat, sisa_berat, sisa int
	fmt.Print("Berat parsel (gram) : ")
	fmt.Scan(&berat_parsel)
	total_berat = berat_parsel / 1000
	sisa_berat = berat_parsel % 1000
	fmt.Printf("Detail Berat : %d kg + %d gr\n", total_berat, sisa_berat)
	detail_biaya = total_berat * 10000
	if sisa_berat >= 500 {
		sisa = sisa_berat * 5
	} else if sisa_berat < 500 {
		sisa = sisa_berat * 15
	} else if sisa_berat > 10 {
		sisa = 0
	}
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d \n", detail_biaya, sisa)
	fmt.Printf("Total biaya: Rp. %d \n", detail_biaya+sisa)
}
