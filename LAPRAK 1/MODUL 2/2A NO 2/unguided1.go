package main

//ARIEL AHNAF KUSUMA 103112400050
import "fmt"

func main() {
	var tahun int
	var kabisat bool
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&tahun)
	kabisat = (tahun%4 == 0 && tahun%100 != 0)
	fmt.Println("Tahun Kabisat:", kabisat)
}
