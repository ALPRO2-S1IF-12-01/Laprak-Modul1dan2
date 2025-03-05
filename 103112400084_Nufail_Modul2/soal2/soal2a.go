package main

// 103112400084 NUFAIL ALAUDDIN TSAQIF
import "fmt"

func main() {
    var k float64

    fmt.Print("Nilai K: ")
    fmt.Scan(&k)

    pembilang := (4*k + 2) * (4*k + 2)
    penyebut := (4*k + 1) * (4*k + 3)
    hasil := pembilang / penyebut

    fmt.Printf("Nilai f(K) = %.10f\n",hasil)
}