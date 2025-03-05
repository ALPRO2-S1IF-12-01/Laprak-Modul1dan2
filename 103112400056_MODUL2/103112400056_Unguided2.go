// Felix Pedrosa V

package main

import "fmt"

func main() {
	var nilaik float64

	for {
		fmt.Print("Nilai K = ")
		_, err := fmt.Scanln(&nilaik)
		if err == nil {
			break
		}
	}

	f_k := (4*nilaik + 2) * (4*nilaik + 2) / ((4*nilaik + 1) * (4*nilaik + 3))

	fmt.Printf("Nilai f(K) = %.10f\n", f_k)
}