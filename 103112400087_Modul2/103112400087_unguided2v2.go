package main

import "fmt"

func main() {
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	hasil := 1.0
	for i := 0; i < k; i++ {
		hasil *= float64((4*i+2)*(4*i+2)) / float64((4*i+1)*(4*i+3))
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}
