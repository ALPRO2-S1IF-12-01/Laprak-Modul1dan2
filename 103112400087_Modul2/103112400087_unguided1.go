package main

import "fmt"

func isLeapYear(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} else if tahun%4 == 0 && tahun%100 != 0 {
		return true
	}
	return false
}
func main() {
	var tahun int
	fmt.Print("masukkan tahun: ")
	fmt.Scan(&tahun)

	if isLeapYear(tahun) {
		fmt.Println("kabisat: true")
	} else {
		fmt.Println("kabisat: false")
	}
}
