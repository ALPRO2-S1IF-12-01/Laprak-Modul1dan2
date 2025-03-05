// MULIA AKBAR NANDA PRATAMA
// 103112400034
package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("masukkan tahun: ")
	fmt.Scan(&tahun)

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Printf("tahun %d merupakan tahun kabisat dengan banyaknya hari adalah 366 (true).\n", tahun)
	} else {
		fmt.Printf("tahun %d merupakan tahun biasa dengan banyaknya hari adalah 365 (false).\n", tahun)
	}
}
