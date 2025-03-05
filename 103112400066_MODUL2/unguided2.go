// DWI OKTA SURYANINGRU
package main

import "fmt"
import "math"

func main() {
	var k, fk float64 // deklarasi variabel k, fk dengan tipe data float64
	
	fmt.Scan(&k) // membaca input dari pengguna dan menyimpan di variabel k

	fk = math.Pow((4*k+2),2) / ((4*k+1) * (4*k+3)) // menghitung (4k + 2) pangkat 2. kemudian dibagi dengan (4k+1) dikali (4k+3)

	fmt.Println("Nilai K =", k) // menampilkan nilai k yang diinputkan user
	fmt.Printf("Nilai f(K) = %.10f\n", fk) // menampilkan hasil perhitungan dari variabel fk dengan presisi 10 angka di belakang koma.
}