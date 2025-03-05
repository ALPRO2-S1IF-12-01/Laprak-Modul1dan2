//Muhammad Zaky Mubarok
package main

import "fmt"

func main() {

var K int

fmt.Print("Nilai K = ")
fmt.Scan(&K)



akar2 := 1.0



for k := 0; k <= K; k++ {

	akar2 *= ((4*float64(k) + 2) * (4*float64(k) + 2)) /

		((4*float64(k) + 1) * (4*float64(k) + 3))

}



fmt.Printf("Nilai akar 2 = %.10f\n", akar2)

}
