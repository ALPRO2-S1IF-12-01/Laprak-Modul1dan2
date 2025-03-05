// DWI OKTA SURYANINGRUM

package main

import "fmt"

func main() {
	var greetings = "Selamat datang di dunia DAB" //membuat variabel greetings dengan tipe data string yang menyimpan teks "Selamat datang di dunia DAB"
	var a, b int //membuat variabel a,b dengan tipe data integer
	
	fmt.Println(greetings) //menampilkan teks yang disimpan di variabel greetings
	fmt.Scan(&a, &b) //membaca input dari pengguna kemudian disimpan ke variabel a dan b

	fmt.Printf("%v + %v = %v\n", a, b, a+b) //menampilkan hasil penjumlahan dalam format. %v digunakan sebagai placeholder untuk menampilkan nilai dari variabel.
	// a + b akan dihitung, lalu hasilnya ditampilkan.
}