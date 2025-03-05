// package main

// import "fmt"

// func main() {
// 	var tahun int
// 	var kabisat bool

// 	fmt.Scan(&tahun)
	
// 	if tahun % 400 == 0 || tahun % 4 == 0 && 4 % 100 != 0{
// 		fmt.Println("Tahun :", tahun)
// 		fmt.Println("Kabisat :", !kabisat)
// 	}else{
// 		fmt.Println("Tahun :", tahun)
// 		fmt.Println("Kabisat :", kabisat)
// 	}
// }

// DWI OKTA SURYANINGRUM
package main

import "fmt"

func main() {
	var tahun int // deklarasi variabel tahun dengan tipe data integer
	var kabisat bool // deklarasi variabel kabisar dengan tipe data boolean

	fmt.Print("Tahun : ") // menampilkan string "Tahun : "
	fmt.Scan(&tahun) // membaca input dari pengguna dan menyimpannya dalam variabel tahun

	//logika dari var kabisat dimana jika tahun habis dibagi 400 atau habis dibagi 4 tetapi tidak habis ketika dibagi 100 maka hasilnya bernilai "true"
	kabisat = tahun % 400 == 0 || tahun % 4 == 0 && 4 % 100 != 0 

	fmt.Println("Kabisat :", kabisat) // menampilkan hasil pengecekan apakah tahun kabisat atau tidak
}