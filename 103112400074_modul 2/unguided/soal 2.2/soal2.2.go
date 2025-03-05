//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func main() {
	var k int
	fmt.Print("nilai K : ")
	fmt.Scan(&k)

	akar2 := 1.0
	for i := 0; i <= k; i++ {
		n := float64(i)
		pembilang := (4*n + 2) * (4*n + 2)
		penyebut := (4*n + 1) * (4*n + 3)
		akar2 *= pembilang / penyebut
	}

	fmt.Printf("Nilai akar 2 : %.10f\n", akar2)
}
