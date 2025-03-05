//Rizkina Azizah_103112400082
package main

import "fmt"

func main(){
	var n int
	var tahun bool
	fmt.Print("Tahun :")
	fmt.Scan(&n)
	tahun = n % 4 == 0 || n % 400 == 0
	fmt.Print("Kabisat :", tahun)
}