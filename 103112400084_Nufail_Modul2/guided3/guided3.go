package main

import "fmt"

func main() {
	nilai := 80
	pcthaadir := 0.75
	adaTubes := true

	var indeks string
	if nilai > 75 && adaTubes {
		indeks = "A"
	} else if nilai > 65 {
		indeks = "B"
	} else if nilai > 55 && pcthaadir > 0.7 {
		indeks = "C"
	} else {
		indeks = "F"
	}
	fmt.Printf("nilai %d dengan kehadiran %.f%% dan buat tubes+%t, mendapat indeks %s\n", nilai, pcthaadir*100, adaTubes, indeks)
}
