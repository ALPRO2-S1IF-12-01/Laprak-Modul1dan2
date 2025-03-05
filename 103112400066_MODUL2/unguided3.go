// DWI OKTA SURYANINGRUM
package main

import "fmt"

func main() {
	var berat_parsel, tambahan_biaya int // deklarasi variabel dengan tipe data integer

	fmt.Print("Berat parsel (gram) : ") // menampilkan teks "Berat parsel (gram) : "
	fmt.Scan(&berat_parsel) // membaca input dari pengguna kemudian menyimpannya dalam variabel berat_parsel

	total_berat := berat_parsel/1000 // perhitungan total berat dalam satuan kilogram (dibagi 1000)
	sisa_berat := berat_parsel - (total_berat*1000) // perhitungan sisa  berat dalam satuan gram (berat yang tidak mencapai 1kg)
	jasa_kirim := total_berat*10000 // biaya pengiriman per kg yaitu 10.000

	if sisa_berat >= 500{ // Jika sisa berat lebih dari atau sama dengan 500 gram, biaya tambahan = Rp. 5 per gram.
		tambahan_biaya = (5*sisa_berat)
	}else{ // Jika sisa berat kurang dari 500 gram, biaya tambahan = Rp. 15 per gram.
		tambahan_biaya =  (15 * sisa_berat)
	}

	total_biaya := jasa_kirim + tambahan_biaya // untuk total biaya dihitung dari jasa kirim + tambahan biaya

	fmt.Println("Detail berat :", total_berat, "kg + ", sisa_berat, "gr") // menampilkan Detail berat dalam kg dan gram.
	fmt.Println("Detail Biaya : Rp.", jasa_kirim, "+ Rp.", tambahan_biaya) // menampilkan Rincian biaya jasa kirim dan tambahan biaya.
	fmt.Println("Total biaya : Rp.", total_biaya) // menampilkan Total biaya pengiriman.
}