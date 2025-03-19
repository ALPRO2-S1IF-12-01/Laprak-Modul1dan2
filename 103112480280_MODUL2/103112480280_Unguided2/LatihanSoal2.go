//Nama : Anggun Wahyu W. (103112480280)
package main

import "fmt"

func main() {
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	FK := func(k int) float64 {
		atas := float64((4*k + 2) * (4*k + 2)) 
		bawah := float64((4*k + 1) * (4*k + 3))
		return atas / bawah
	}

	akar2 := func(k int) float64 {
		hasil := 1.0
		for i := 0; i <= k; i++ {
			atas := float64((4*i + 2) * (4*i + 2)) 
			bawah := float64((4*i + 1) * (4*i + 3))
			hasil *= atas / bawah
		}
		return hasil
	}

	fmt.Printf("Nilai f(K) = %.10f\n", FK(k))
	fmt.Printf("Nilai akar 2 = %.10f\n", akar2(k))
}
