package main

import "fmt"

func main() {
	var t int
	var ket bool
	fmt.Scan(&t)

	if t % 400 == 0 || t % 4 == 0 {
		ket = true
	} else if t % 100 != 0 {
		ket = false 
	}

	fmt.Print(ket)
}
