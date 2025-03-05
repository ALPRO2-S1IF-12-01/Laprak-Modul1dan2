// RAJA MUHAMMAD LUFHTI || 103112400027 
package main

import "fmt"

func main() {
	var q int
	var root2 float64
	fmt.Scan(&q)
	root2 = 1.0
	for i := 0; i <= q; i++ {
		root2 *= float64((4*i+2)*(4*i+2)) / float64((4*i+1)*(4*i+3))
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", root2)
}
