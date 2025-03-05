//ICHYA ULUMIDDIIN
package main

import (
	"fmt"
)
func main() {
	var beratParsel int
	fmt.Print("Berat Parsel (gram): ")
	fmt.Scan(&beratParsel)
	beratKg := beratParsel / 1000
	sisaGr := beratParsel % 1000
	biaya := beratKg * 10000
	if beratKg > 10 {
		sisaGr = 0
	}
	if sisaGr >= 500 {
		biaya += sisaGr * 5
	} else {
		biaya += sisaGr * 15
	}
	fmt.Printf("\nDetail berat : %d kg + %d gr\n", beratKg, sisaGr)
	fmt.Printf("Detail Biaya : Rp. %d + Rp. %d\n", beratKg*10000, biaya-(beratKg*10000))
	fmt.Printf("Total Biaya : Rp. %d\n", biaya)
}
