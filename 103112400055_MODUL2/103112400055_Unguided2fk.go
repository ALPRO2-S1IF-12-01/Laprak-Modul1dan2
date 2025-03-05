//Feros Pedrosa

package main

import "fmt"

func main() {
	var k float64

	for {
		fmt.Print("Nilai K = ")
		_, err := fmt.Scanln(&k)
		if err == nil {
			break
		}
		fmt.Println("Input tidak valid.")
	}

	f_k := (4*k + 2) * (4*k + 2) / ((4*k + 1) * (4*k + 3))

	fmt.Printf("Nilai f(K) = %.10f\n", f_k)
}
