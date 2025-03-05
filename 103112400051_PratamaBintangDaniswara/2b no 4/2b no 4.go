// PRATAMA BINTANG DANISWARA
// 103112400051
package main

import "fmt"

func main() {
	var K int
	fmt.Print("Masukan nilai K: ")
	fmt.Scan(&K)
	atas := float64((4*K + 2) * (4*K + 2))
	bawah := float64((4*K + 1) * (4*K + 3))
	hasil := float64((atas / bawah))
	fmt.Printf("Nilai f(k) = %.10f\n", hasil)
}
