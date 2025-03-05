//Nama	: Raihan Adi Arba
//NIM	: 103112400071
//KELAS	: IF-12-01

package main

import "fmt"

func main() {
	var tahun int
	var ckabisat bool

	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	if tahun%100 == 0 {
		ckabisat = tahun%400 == 0
	} else {
		ckabisat = tahun%4 == 0
	}

	fmt.Printf("Kabisat: %t\n", ckabisat)
}
