//M.DAVI ILYAS RENALDO
//103112400062

package main

import "fmt"

func akar2(K int) float64 {
        hasil := 1.0
        for k := 0; k <= K; k++ {
                pembilang := float64((4*k + 2) * (4*k + 2))
                penyebut := float64((4*k + 1) * (4*k + 3))
                hasil *= pembilang / penyebut
        }
        return hasil
}

func main() {
        var K int
        fmt.Print("Nilai K = ")
        fmt.Scan(&K)

        hasil := akar2(K)
        fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}