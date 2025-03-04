package main

import (
	"fmt"
	"math"
)

func hitungAkar2(K int) float64 {
	akar2 := 1.0
	for k := 0; k <= K; k++ {
		akar2 *= math.Pow(float64(4*k+2), 2) / float64((4*k+1)*(4*k+3))
	}
	return akar2
}

func main() {
	var K int
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&K)
	fmt.Printf("Akar 2 = %.10f\n", hitungAkar2(K))
}
