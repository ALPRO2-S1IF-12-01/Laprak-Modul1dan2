package main

import "fmt"

func main() {
	var k int
	var fK, hasil, term float64
	hasil = 1

	fmt.Print("Nilai K: ")
	fmt.Scan(&k)

	for i := 0; i <= k; i++ {
		term = float64((4*i+2)*(4*i+2)) / float64((4*i+1)*(4*i+3))
		hasil *= term
		if i == k {
			fK = term
		}
	}

	fmt.Printf("Nilai f(K) = %.10f\n", fK)
	fmt.Printf("Nilai hampiran \u221A2 = %.10f\n", hasil)
}
