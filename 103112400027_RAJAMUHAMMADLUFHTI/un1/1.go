// RAJA MUHAMMAD LUFHTI || 103112400027 
package main

import "fmt"

func main() {
	var O int
	var hasil bool
	fmt.Scan(&O)
	hasil =  true

	if O%400 == 0 || O%4 == 0 && O%100 != 0 {
		fmt.Println("year :", O)
		fmt.Println("leap :", hasil)
		
	}else {
		fmt.Println("year :", O)
		fmt.Println("leap :", !hasil)
	}
}