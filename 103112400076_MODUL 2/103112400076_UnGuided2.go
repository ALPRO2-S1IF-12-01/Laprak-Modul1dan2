package main

import "fmt"

func main() {
	var k int
	fmt.Print("Masukkan Nilai K: ")
	fmt.Scan(&k)
	hasil := 1.0
	for i := 0; i <= k; i++ {
		penyebutAtas := (4*float64(i) + 2) * (4*float64(i) + 2)
		penyebutBawah := (4*float64(i) + 1) * (4*float64(i) + 3)
		hasil *= penyebutAtas / penyebutBawah
	}
	fmt.Printf("%.10f\n", hasil)
}
