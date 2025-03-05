//Rizkina Aziazah_103112400082
package main

import "fmt"

func main(){
	var k int
	var akar2 float64
	fmt.Print("k: ")
	fmt.Scan(&k)
	akar2 = 1.0
	for i := 0; i <= k; i++ {
		akar2 *= (float64(4*i + 2)) * (float64(4*i + 2)) / (float64(4*i + 1) * float64(4*i + 3))
	
	}
	fmt.Printf("akar2: %.10f\n", akar2)
}