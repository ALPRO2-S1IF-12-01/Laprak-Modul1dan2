package main

import (
	"fmt"
	"math"
)

func hitungAkar2(K int) float64 {
	var hasil float64 = 1.0

	for k := 0; k <= K; k++ {
		nom := math.Pow(float64(4*k+2), 2)
		den := float64((4*k + 1) * (4*k + 3))
		hasil *= nom / den
	}

	akar2 := hasil
	return akar2
}

func main() {
	var K int
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&K)
	akar2 := hitungAkar2(K)
	fmt.Printf("Akar 2 = %.10f\n", akar2)
}
