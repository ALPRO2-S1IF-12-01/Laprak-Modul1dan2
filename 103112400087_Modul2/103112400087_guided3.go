//103112400087

package main

import "fmt"

func main() {
	nilai := 80
	pcthadir := 0.75
	adatubes := true

	var indeks string
	if nilai > 75 && adatubes {
		indeks = "A"
	} else if nilai > 65 {
		indeks = "B"
	} else if nilai > 50 && pcthadir > 0.7 {

		indeks = "C"
	} else {
		indeks = "f"
	}
	fmt.Printf("nilai %d dengan kehadiran %.2f%% dan buat tubes=%t, mendapatkan indeks %s\n", nilai, pcthadir*100, adatubes, indeks)
}
