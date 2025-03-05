package main

import "fmt"

func main() {
	var berat, kg, gr, biaya, tambahan int
	fmt.Scan(&berat)

	kg = berat / 1000
	gr = berat % 1000
	biaya = kg * 10000

	if gr >= 500 {
		tambahan = gr * 5
	} else if gr < 500 {
		tambahan = gr * 15
	}

	total := biaya + tambahan

	if kg > 10 {
		fmt.Println("Berat parsel (gram): ", berat)
		fmt.Printf("Detail berat : %v kg + %v gr\n", kg, gr)
		fmt.Printf("Detail biaya: Rp. %v + Rp. %v\n", biaya, tambahan)
		fmt.Printf("Total biaya: Rp. %v", biaya)
	} else if kg < 10 {
		fmt.Println("Berat parsel (gram): ", berat)
		fmt.Printf("Detail berat : %v kg + %v gr\n", kg, gr)
		fmt.Printf("Detail biaya: Rp. %v + Rp. %v\n", biaya, tambahan)
		fmt.Printf("Total biaya: Rp. %v", total)
	}
}
