//103112400070 Achmad Zulvan Nur Hakim

package main

import "fmt"

func main() {
	var weightGr, addcost int
	fmt.Print("Berat(gram): ")
	fmt.Scanln(&weightGr)

	weightKg := weightGr / 1000
	remainingGr := weightGr % 1000
	cost := weightKg * 10000

	fmt.Printf("Detail Berat: %d kg + %d gr\n", weightKg, remainingGr)

	
	if weightKg > 10 {
		fmt.Printf("Total : Rp. %d\n", cost)
		return
	}
	
	if remainingGr >= 500 {
		addcost = remainingGr * 5
	} else {
		addcost = remainingGr * 15
	}

	fmt.Printf("Detail biaya: Rp. %d + Rp. %d \n", cost, addcost)
	fmt.Printf("Total : Rp. %d\n", cost+addcost)
}
