//ABID FADHILAH M
//103112400046
package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("masukan tahun: ")
	fmt.Scan(&tahun)
	if(tahun%4 == 0 && tahun%100 !=0) || (tahun%400 == 0){
		fmt.Printf("tahun %d adalah kabisat dengan 366 hari (true).\n", tahun)
	}else {
		fmt.Printf("tahun %d adalah tahun biasa dengan 365 hari (false).\n, tahun")
	}
}