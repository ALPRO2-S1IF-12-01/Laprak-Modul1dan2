// 103112400087

package main

import "fmt"

func main() {
	var k float64
	fmt.Print("Nilai k = ")
	fmt.Scan(&k)

	fk := ((4*k + 2) * (4*k + 2)) / ((4*k + 1) * (4*k + 3))
	fmt.Printf("nilai f(k) = %.10f\n", fk)
}
