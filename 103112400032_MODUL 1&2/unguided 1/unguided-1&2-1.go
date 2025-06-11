// daffa tsaqifna f
package main

import "fmt"

func kabisat(tahun int) bool {
	return (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)
}
func main() {
	var x int
	fmt.Print("tahun: ")
	fmt.Scan(&x)
	fmt.Print("kabisat: ", kabisat(x))
}
