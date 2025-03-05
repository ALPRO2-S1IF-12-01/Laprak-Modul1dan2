// M.HANIF AL FAIZ
// 103112400042
package main

import "fmt"

func main() {
	var tahun int
	var status bool
	fmt.Scan(&tahun)
	status = tahun%4 == 0
	fmt.Println("TAHUN:", tahun)
	fmt.Println("KABISAT:", status)
}
