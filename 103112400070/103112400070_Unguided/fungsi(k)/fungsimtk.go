//103112400070 Achmad Zulvan Nur Hakim
package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	fmt.Scan(&k)

	fk := 1.0
	for i := 0; i < k; i++ {
		fk *= math.Pow(float64((4*i+2)), 2) / float64((4*i+1)*(4*i+3))
	}

	fmt.Printf("Nilai f(k) = %.10f\n", fk)
}
