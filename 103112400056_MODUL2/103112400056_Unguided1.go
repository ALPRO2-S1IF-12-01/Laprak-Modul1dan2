// Felix Pedrosa V

package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	if tahun%400 == 0 {
		fmt.Printf("Kabisat: true\n")
	} else if tahun%100 == 0 {
		fmt.Printf("Kabisat: false\n")
	} else if tahun%4 == 0 {
		fmt.Printf("Kabisat: true\n")
	} else {
		fmt.Printf("Kabisat: false\n")
	}
}
