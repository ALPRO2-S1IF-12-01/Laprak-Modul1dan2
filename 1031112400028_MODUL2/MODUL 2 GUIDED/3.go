package main
import "fmt"
func main() {
	nilai := 80
	pchadir := 0.75
	adatubes := true

	var indeks string

	if nilai > 75 && adatubes {
		indeks = "A"
	} else if nilai > 65 {
		indeks = "B"
	} else if nilai > 50 && pchadir > 0.7 {
		indeks = "C"
	} else {
		indeks = "f"
	}
	fmt.Printf("nilai %d dengan kehadiran %.2f%% dan buat tubes=%t, mendapat indeks %s\nilai", pchadir*100, adatubes, indeks)
}