// daffa tsaqifna f
package main

import "fmt"

func biaya(rp int) int {
	var x, y int
	x = rp / 1000
	x *= 10000
	if rp%1000 != 0 {
		y = rp - (x / 10)
		if y >= 500 {
			y *= 5
			return x + y
		} else {
			y *= 15
			return x + y
		}
	} else {
		return x
	}
}
func main() {
	var x, y int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&x)
	y = x % 1000
	if y >= 500 {
		y *= 5
	} else {
		y *= 15
	}
	fmt.Printf("Detail berat: %d kg + %d gr\n", x/1000, x%1000)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", (x/1000)*10000, y)
	fmt.Printf("Total biaya: Rp. %d\n", biaya(x))
}
