// Muhamad Faza Fahri Aziz || 103112400072

package main

import (
	"fmt"
	"math"
)

type kalkulatorAkar2 struct{}

func (c kalkulatorAkar2) hitung(K int) float64 {
	result := 1.0
	for k := 0; k <= K; k++ {
		result *= math.Pow(float64(4*k+2), 2) / float64((4*k+1)*(4*k+3))
	}
	return result
}

func main() {
	var K int
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&K)

	calc := kalkulatorAkar2{}
	fmt.Printf("Akar 2 = %.10f\n", calc.hitung(K))
}
