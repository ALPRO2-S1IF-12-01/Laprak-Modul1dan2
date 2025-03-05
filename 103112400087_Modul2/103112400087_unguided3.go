package main

import "fmt"

func main() {
	var berat int
	fmt.Print("berat parsel (gram): ")
	fmt.Scan(&berat)
	kg := berat / 1000
	gram := berat % 1000
	biaya := kg * 10000
	if berat > 10000 {
		gram = 0
	}
	if gram >= 500 {
		biaya += gram * 5
	} else {
		biaya += gram * 15
	}

	fmt.Printf("detail berat: %d kg + %d gr\n", kg, gram)
	fmt.Printf("total biaya: Rp. %d\n", biaya)
}
