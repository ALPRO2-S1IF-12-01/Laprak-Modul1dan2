package main

//103112400084 NUFAIL ALAUDDIN TSAQIF
import "fmt"

func main() {
	var k int
	var hasil float64
	
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	hasil = 1.0

	for i := 0; i <= k; i++ {
		pembilang := (4*i + 2) * (4*i + 2)
		penyebut := (4*i + 1) * (4*i + 3)
		hasil = hasil * float64(pembilang) / float64(penyebut)
	}
	fmt.Printf("nilai akar 2 = %.10f\n", hasil)
}
