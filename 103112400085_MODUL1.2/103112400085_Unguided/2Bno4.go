package main

import "fmt"

func hitungFk(K int) float64 {

	hasil := 1.0
	for i := 0; i < K; i++ {
		penyebut := (4*float64(i) + 2) * (4*float64(i) + 2)
		pembilang := (4*float64(i) + 1) * (4*float64(i) + 3)

		hasil *= float64(penyebut) / float64(pembilang)
	}
	return hasil
}

func main() {
	var K int
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	akar := hitungFk(K)
	fmt.Printf("Nilai akar 2 = %.10f\n", akar)
}
