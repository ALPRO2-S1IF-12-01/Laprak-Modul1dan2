//103112400070 Achmad Zulvan Nur Hakim
package main

import "fmt"

func main() {
	var th int
	fmt.Scan(&th)
	if th%400 == 0 || (th%4 == 0 && th%100 != 0) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")

	}
}
