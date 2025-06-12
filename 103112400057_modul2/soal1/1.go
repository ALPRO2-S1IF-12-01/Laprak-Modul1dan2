package main

import (
	"fmt"
	"strings"
)

func tampilkanGaris() {
	fmt.Println(strings.Repeat("=", 50))
}
func cekTahunKabisat(tahun int) bool {
	return (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
	var tahun int
	tampilkanGaris()
	fmt.Println("PROGRAM CEK TAHUN KABISAT")
	tampilkanGaris()
	fmt.Print("\nMasukkan tahun: ")
	fmt.Scan(&tahun)
	if tahun <= 0 {
		fmt.Println("\nError: Tahun harus lebih besar dari 0!")
		return
	}
	tampilkanGaris()
	fmt.Println("HASIL PENGECEKAN")
	tampilkanGaris()

	if cekTahunKabisat(tahun) {
		fmt.Printf("Tahun %d adalah tahun kabisat\n", tahun)
		fmt.Println("Jumlah hari: 366 hari")
		fmt.Println("Status: True")
	} else {
		fmt.Printf("Tahun %d adalah tahun biasa\n", tahun)
		fmt.Println("Jumlah hari: 365 hari")
		fmt.Println("Status: False")
	}
	tampilkanGaris()
}
