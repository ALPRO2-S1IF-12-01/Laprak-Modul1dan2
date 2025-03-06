package main

//ARIEL AHNAF KUSUMA 103112400050
import "fmt"

func main() {
	var beratgram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratgram)

	kg, gram := beratgram/1000, beratgram%1000
	biayakg, biayagram := kg*10000, 0

	if gram >= 500 {
		biayagram = gram * 5
	} else {
		biayagram = gram * 15
	}

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayakg, biayagram)
	fmt.Printf("Total biaya: Rp. %d\n", biayakg+biayagram)
}
