package main

import "fmt"

func main() {
	var k int
	var hasil float64
	fmt.Print("Nilai K : ")
	fmt.Scan(&k)

	for i := 0; i < k; i++ {
		pembilang := (4*i + 2) * (4*i + 2)
		penyebut := (4*i + 1) * (4*i + 3)

		if i == 0 {
			hasil = float64(pembilang) / float64(penyebut)
		} else {
			hasil = hasil * (float64(pembilang) / float64(penyebut))
		}
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}
