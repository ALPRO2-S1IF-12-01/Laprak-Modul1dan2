package main

import "fmt"

func kabisat(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} else if tahun%100 == 0 {
		return false
	} else if tahun%4 == 0 {
		return true
	}
	return false
}

func main() {
	var tahun int
	fmt.Print("Masukan Tahun: ")
	fmt.Scanln(&tahun)

	if kabisat(tahun) {
		fmt.Println("Kabisat : true")
	} else {
		fmt.Println("Kabisat : false")
	}
}
