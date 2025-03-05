// HAKAN ISMAIL AFNAN
// 103112400038
package main

import (
	"fmt"
)

func main() {
	var tahun int
	fmt.Print("Masukkan tahun yang ingin diperiksa: ")
	fmt.Scanln(&tahun)
	switch {
	case (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0):
		fmt.Printf("%d merupakan tahun kabisat dengan 366 hari.\n", tahun)
	default:
		fmt.Printf("%d bukan tahun kabisat, hanya memiliki 365 hari.\n", tahun)
	}
}
