// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func main() {
	var tahun int
	fmt.Scan(&tahun)

	kabisat := tahun % 400 == 0 || (tahun % 4 == 0 && tahun % 100 != 0)

	if kabisat {
		fmt.Print("Kabisat: true")
	} else {
		fmt.Print("Kabisat: false")
	}
}