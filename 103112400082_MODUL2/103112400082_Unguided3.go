//Rizkina Azizah_10312400082
package main
import "fmt"

func main(){
	var  bp, tambah int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&bp)
	kg := bp/1000
	sisa_gr := bp%1000
	kirim := kg * 10000
	if kg > 10 {
		tambah = 0
	}else if sisa_gr >= 500 {
		tambah = sisa_gr * 5
	}else if sisa_gr< 500 {
		tambah = sisa_gr * 15
	}
	total := kirim + tambah
	fmt.Printf("Detail Berat: %d kg + %d gr\n ", kg, sisa_gr)
	fmt.Printf("Detail Biaya: Rp. %d kg + Rp. %d\n", kirim, tambah)
	fmt.Printf("Total biaya: Rp. %d\n ", total)
}

