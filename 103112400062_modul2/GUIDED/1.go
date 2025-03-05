package main

import "fmt"

func main() {
	var greetings = "selamat datang didunia DAP"
	var a, b int
	fmt.Println(greetings)
	fmt.Scan(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)

}