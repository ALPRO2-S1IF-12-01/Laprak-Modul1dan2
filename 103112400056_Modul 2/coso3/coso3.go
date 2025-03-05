// Felix Pedrosa V

package main

import "fmt"
func main() {
	// Contoh nilai (bisa diganti sesuai kebutuhan)
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

	fmt.Printf("Nilai %d dengan kehadiran %.2f%% dan buat tubes = %t memiliki indeks %s\n", nilai, pctHadir*100, adaTubes, indeks)
}