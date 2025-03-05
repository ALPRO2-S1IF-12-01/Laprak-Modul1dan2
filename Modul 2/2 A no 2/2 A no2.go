//Muhammad Zaky Mubarok
package main

import (
    "fmt"
)

func main() {
    var tahun int
    fmt.Print("Tahun: ")
    fmt.Scan(&tahun)

    kabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

    if kabisat {
        fmt.Println("Kabisat: true")
    } else {
        fmt.Println("Kabisat: false")
    }
}