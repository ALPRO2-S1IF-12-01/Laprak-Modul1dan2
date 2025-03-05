// Ahmad Ruba'i
//NIM: 103112400074
package main

import "fmt"

func main() {
	var greetings = "selamat datang di dunia DAP"
	var a, b int
	fmt.Println(greetings)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
