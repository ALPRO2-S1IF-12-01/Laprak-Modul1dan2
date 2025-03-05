// Ahmad Ruba'i
// 103112400074
package main

import "fmt"

func main() {
	var tahun int
	fmt.Scan(&tahun)

	kabisat := tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)

	fmt.Printf("Tahun: %d\nKabisat: %t\n", tahun, kabisat)
}
