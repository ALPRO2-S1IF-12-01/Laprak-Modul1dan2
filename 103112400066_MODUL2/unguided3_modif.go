// DWI OKTA SURYANINGRUM
package main

import (
	"fmt"
	"math"
)

func main() {
	var k int // deklarasi variabel dengan tipe data integer

	fmt.Print("Nilai K: ") // menampilkan string "Nilai K: "
	fmt.Scan(&k) // membaca input dari user dan menyimpan di variabel k

	result := 1.0 // deklarasi vairabel result dengan tipe data float64 dan menyimpan nilai 1.0

	for i := 0; i < k; i++ { // (perulangan) melakukan perhitungan sebanyak K kali:
		numerator := math.Pow(float64(4*i+2), 2) // Menghitung pangkat dua dari (4*i + 2)
		denumerator := float64((4*i + 1) * (4*i + 3)) //Mengalikan (4*i + 1) dengan (4*i + 3)
		result *= numerator / denumerator // Hasil bagi pembilang dan penyebut akan dikalikan dengan variabel result secara terus-menerus.
	}

	fmt.Printf("Hasil dari operasi fungsi = %.10f\n", result) // hasil akhir ditampilkan dengan ketelitian 10 angka di belakang koma:
}