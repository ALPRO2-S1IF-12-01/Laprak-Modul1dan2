//ICHYA ULUMIDDIIN
package main

import "fmt"

func main() {
    var tahun int
	fmt.Print("Tahun: ")
    fmt.Scan(&tahun)
    fmt.Println("Kabisat :", (tahun % 400 == 0) || (tahun % 4 == 0 && tahun % 100 != 0))
}
