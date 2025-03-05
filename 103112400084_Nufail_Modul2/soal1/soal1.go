package main
//103112400084 NUFAIL ALAUDDIN TSAQIF
import "fmt"
func main() {
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)
	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println("kabisat: true")
	} else {
		fmt.Println("kabisat: false")
	}
}