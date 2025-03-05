// Muhamad Faza Fahri Aziz || 103112400072 || Alpro 2

package main

import "fmt"

func main() {
	var n int
	var hasil bool
	fmt.Scan(&n)
	hasil =  true

	if n%400 == 0 || n%4 == 0 && n%100 != 0 {
		fmt.Println("Tahun :", n)
		fmt.Println("Kabisat :", hasil)
		
	}else {
		fmt.Println("Tahun :", n)
		fmt.Println("Kabisat :", !hasil)
	}
}