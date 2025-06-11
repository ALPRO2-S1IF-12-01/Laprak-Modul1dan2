// daffa tsaqifna f
package main

import "fmt"

func fx(k int) float64 {
	var x float64
	x = 1.0
	for i := 0; i < k; i++ {
		y := float64((4*i + 2) * (4*i + 2))
		z := float64((4*i + 1) * (4*i + 3))
		x *= y / z
	}
	return x
}
func main() {
	var x int
	fmt.Print("nilai K = ")
	fmt.Scan(&x)
	fmt.Printf("nilai f(k) = %.10f", fx(x))
}
