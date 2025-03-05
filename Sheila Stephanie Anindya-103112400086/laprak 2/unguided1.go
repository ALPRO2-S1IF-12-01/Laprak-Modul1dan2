package main

import "fmt"

func main() {
	var tahun int
	fmt.Scan(&tahun)

	if tahun%4 == 0 && tahun%100 != 0 {
		fmt.Println(true)
	} else if tahun%400 == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

// Sheila Stephanie Anindya - 10311240086
