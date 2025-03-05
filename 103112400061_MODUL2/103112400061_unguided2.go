package main

import "fmt"

func main() {
	var (
		K, toK int
		sqrt2  float64 = 1.0
	)
	fmt.Print("Nilai K = ")
	fmt.Scan(&toK)
	for K = 0; K <= toK; K++ {
		pembilang := float64((4*K + 2) * (4*K + 2))
		penyebut := float64((4*K + 1) * (4*K + 3))
		sqrt2 *= pembilang / penyebut
	}
	fmt.Printf("Nilai akar 2 = %.10f \n", sqrt2)
}