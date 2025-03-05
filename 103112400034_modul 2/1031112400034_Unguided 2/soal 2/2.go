// MULIA AKBAR NANDA PRATAMA
// 103112400034
package main

import "fmt"

func main() {
	var K int

	fmt.Print("masukkan nilai K = ")
	fmt.Scan(&K)

	akar2 := 1.0

	for k := 0; k <= K; k++ {
		akar2 *= float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}
