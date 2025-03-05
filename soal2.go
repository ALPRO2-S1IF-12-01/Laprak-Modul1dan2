//ABID FADHILAH M
//103112400046
package main

import (
    "fmt"
    "math"
)

func kalkulasiAkar2(A int) float64 {
    hasilA := 1.0
    for a := 0; a <= A; a++ {
        hasilA *= math.Pow(float64(4*a+2), 2) / float64((4*a+1)*(4*a+3))
    }
    return hasilA
}

func main() {
    var A int
    fmt.Print("Masukkan nilai A: ")
    fmt.Scan(&A)
    fmt.Printf("Akar 2 = %.10f\n", kalkulasiAkar2(A))
}

