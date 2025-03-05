package main

import "fmt"

func main() {
	var greetings = "selamat datang di dunia DAP"
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(greetings)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}