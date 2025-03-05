//Nama	: Raihan Adi Arba
//NIM	: 103112400071
//KELAS	: IF-12-01

package main

import "fmt"

func main() {
	nilai := 80
	pctHadir := 0.75
	adaTubes := true

	var indeks string

	if nilai > 75 && adaTubes {
		indeks = "A"
	} else if nilai > 65 {
		indeks = "B"
	} else if nilai > 50 && pctHadir > 0.7 {
		indeks = "C"
	} else {
		indeks = "F"
	}
	fmt.Printf("nilai %d dengan kehadiran %.2f%% dan buat tubes=%t, mendapat indeks %s\n", nilai, pctHadir*100, adaTubes, indeks)
}
