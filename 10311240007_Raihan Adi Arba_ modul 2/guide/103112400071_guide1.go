//Nama	: Raihan Adi Arba
//NIM	: 103112400071
//KELAS	: IF-12-01

package main

import "fmt"

func main() {
	var greeting string = "Selamat datang di dunia DAP"
	var a, b int
	fmt.Println(greeting)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
