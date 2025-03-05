package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&K)

	akr2 := 1.0

	for k := 0; k <= K; k++ {
		akr2 *= ((4*float64(k) + 2) * (4*float64(k) + 2)) /
			((4*float64(k) + 1) * (4*float64(k) + 3))
	}

	fmt.Printf("%.10f\n", akr2)
}
