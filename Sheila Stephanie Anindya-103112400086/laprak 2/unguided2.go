package main

import "fmt"

func main() {
	var i, k int
	var fk float64
	fmt.Print("Nilai K : ")
	fmt.Scan(&k)

	for i = k; i <= k; i++ {
		pembilang := (4*i + 2) * (4*i + 2)
		penyebut := (4*i + 1) * (4*i + 3)
		fk = float64(pembilang) / float64(penyebut)
	}

	fmt.Printf("Nilai f(K) = %.10f\n", fk)
}
