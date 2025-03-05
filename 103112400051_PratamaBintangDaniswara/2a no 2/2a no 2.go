// PRATAMA BINTANG DANISWARA
// 103112400051

package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	switch {
	case tahun%400 == 0:
		fmt.Print("Kabisat: true")

	case tahun%100 == 0:
		fmt.Print("Kabisat: false")

	case tahun%4 == 0:
		fmt.Print("Kabisat: true")

	default:
		fmt.Print("Kabisat: false")
	}
}
