package main

import (
	"fmt"
)

func main() {
	var berat_asli int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat_asli)

	kg := berat_asli / 1000
	gram := berat_asli % 1000

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)

	biaya_kg := kg * 10000
	var biaya_gram int

	if gram >= 500 {
		biaya_gram = gram * 5
	} else {
		biaya_gram = gram * 15
	}

	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biaya_kg, 
	fmt.Printf("Total biaya: Rp. %d", biaya_kg+biaya_gram)
}
