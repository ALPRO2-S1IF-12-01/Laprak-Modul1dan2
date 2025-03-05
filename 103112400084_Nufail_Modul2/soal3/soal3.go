package main
// 103112400084 NUFAIL ALAUDDIN TSAQIF
import "fmt"
func main() {
    var BeratGram int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&BeratGram)
    BeratKg := BeratGram / 1000
    SisaGram := BeratGram % 1000
    BiayaKg := BeratKg * 10000
    BiayaSisa := 0

    if BeratKg > 10 && SisaGram < 1000 {
        BiayaSisa = 0
    } else if SisaGram >= 500 {
        BiayaSisa = SisaGram * 5
    } else {
        BiayaSisa = SisaGram * 15
    }
    totalBiaya := BiayaKg + BiayaSisa
    fmt.Printf("Detail berat: %d kg + %d gr\n", BeratKg, SisaGram)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", BiayaKg, BiayaSisa)
    fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}