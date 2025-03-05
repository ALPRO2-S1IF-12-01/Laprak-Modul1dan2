// DWI OKTA SURYANINGRUM
package main

import "fmt"

func main() {
	nilai := 80 // deklarasi variabel dengan tipe data int dan menyimpan nilai 80
	pctHadir := 0.75 // deklarasi variabel dengan tipe data float64 dan menyimpan nilai 0.75 / 75% sebagai (presentase kehadiran)
	adaTubes := true // deklarasi variabel dengan tipe data boolean bernilai "true" yang menunjukkan mahasiswa mengerjakan tubes

	var indeks string //deklarasi variabel indeks dengan tipe data string

	if nilai > 75 && adaTubes { // kondisi jika nilai lebih dari 75 dan mengerjakan tubes maka indeksnya adalah "A"
		indeks = "A"
	}else if nilai > 65 { // jika nilai lebih dari 65 maka indeksnya "B"
		indeks = "B"
	}else if nilai > 50 && pctHadir > 0.7 { //jika nilai lebih dari 50 dan presentasi hadir lebih dari 70% maka indeksnya "C"
		indeks = "C"
	}else { // tidak memenuhi semua syarat maka indeksnya "F"
		indeks = "F"
	}

	// output akan menampilkan hasil berupa nilai, kehadiran, tugas besar, dan indeks yang didapat.
	fmt.Printf("Nilai %d dengan kehadiran %.2f%% dan buat tubes = %t, mendapatkan indeks %s\n", nilai, pctHadir*100, adaTubes, indeks)
}