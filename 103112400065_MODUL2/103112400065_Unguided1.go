package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)
	if tahun%400 == 0 {
		fmt.Println("Tahun Kabisat")
	} else if tahun%4 == 0 && tahun%100 != 0 {
		fmt.Println("Tahun Kabisat")
	} else {
		fmt.Println("Bukan Tahun Kabisat")
	}
}
