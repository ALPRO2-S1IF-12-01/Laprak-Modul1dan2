// MUHAMMAD FAUZAN
package main

import "fmt"

func main() {
	var tahun int
	fmt.Scan(&tahun)
	if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
		fmt.Printf("Tahun %d adalah tahun kabisat dengan 366 hari (true).\n", tahun)
	} else {
		fmt.Printf("Tahun %d adalah tahun biasa dengan 365 hari (false).\n", tahun)
	}
}
