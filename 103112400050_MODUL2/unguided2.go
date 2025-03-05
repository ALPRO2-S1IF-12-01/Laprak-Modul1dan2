package main

//ARIEL AHNAF KUSUMA 103112400050
import "fmt"

func main() {
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	akar2 := 1.0
	for k := 0; k <= k; k++ {
		x := 4*float64(k) + 2
		y := 4*float64(k) + 1
		akar2 *= (x * x) / (y * (y + 2))
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}
