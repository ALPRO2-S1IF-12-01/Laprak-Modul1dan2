// Felix Pedrosa V

package main

import "fmt"

func main() {
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	hasil := 1.0
	for i := 0; i <= k; i++ {
		pembilang := (4*i + 2) * (4*i + 2)
		penyebut := (4*i + 1) * (4*i + 3)
		hasil *= float64(pembilang) / float64(penyebut)
	}
	fmt.Printf("nilai √2 = %.10f\n", hasil)
}

