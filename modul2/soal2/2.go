package main

import (
	"fmt"
	"strings"
	"math"
)
func tampilkanGaris() {
	fmt.Println(strings.Repeat("=", 50))
}
func hitungAkar2(K int) float64 {
	akar2 := 1.0
	for k := 0; k <= K; k++ {
		akar2 *= float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
	}
	return akar2
}

func main() {
	var K int
	tampilkanGaris()
	fmt.Println("PROGRAM MENGHITUNG NILAI AKAR 2")
	tampilkanGaris()
	fmt.Print("\nMasukkan nilai K (0-100): ")
	fmt.Scan(&K)
	if K < 0 || K > 100 {
		fmt.Println("\nError: Nilai K harus antara 0-100!")
		return
	}
	hasil := hitungAkar2(K)
	nilaiSebenarnya := math.Sqrt(2)
	tampilkanGaris()
	fmt.Println("HASIL PERHITUNGAN")
	tampilkanGaris()
	fmt.Printf("Nilai K\t\t: %d\n", K)
	fmt.Printf("Hasil perhitungan\t: %.10f\n", hasil)
	fmt.Printf("Nilai sebenarnya\t: %.10f\n", nilaiSebenarnya)
	fmt.Printf("Selisih\t\t: %.10f\n", math.Abs(hasil-nilaiSebenarnya))
	tampilkanGaris()
}