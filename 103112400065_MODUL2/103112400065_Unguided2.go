package main

import "fmt"

func main() {
	var hasil, a float64
	var k int
	a = 0
	hasil = 1
	fmt.Print("Nilai K: ")
	fmt.Scan(&k)
	for i := 0; i <= k; i++ {
		temp := (4*a + 2) * (4*a + 2) / (4*a + 1) / (4*a + 3)
		hasil *= temp
		a++
	}
	fmt.Printf("Nilai akar 2 = %.10f", hasil)
}
