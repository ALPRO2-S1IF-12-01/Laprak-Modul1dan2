//ABID FADHILAH M
//103112400046
package main

import (
    "fmt"
)

func hitungbiaya(berat int) int {
    kg := berat / 1000
    gram := berat % 1000
    biayatotal := kg * 10000
    
    if gram >= 500 {
        biayatotal += gram * 5
    } else {
        biayatotal += gram * 15
    }
    return biayatotal
}

func main() {
    var berat int
    fmt.Print("masukkan berat parsel (gram): ")
    fmt.Scanln(&berat)
    fmt.Printf("total biaya: Rp. %d\n", hitungbiaya(berat))
}

