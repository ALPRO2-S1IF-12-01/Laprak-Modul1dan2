//Nama	: Raihan Adi Arba
//NIM	: 103112400071
//KELAS	: IF-12-011

package main

import (
	"fmt"
)

func main() {
	var xy, x, y float64
	var vloop, K int

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)
	xy = 1.0

	for vloop = 0; vloop <= K; vloop++ {
		x = float64((4*vloop + 2) * (4*vloop + 2))
		y = float64((4*vloop + 1) * (4*vloop + 3))
		xy *= x / y
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", xy)
}
