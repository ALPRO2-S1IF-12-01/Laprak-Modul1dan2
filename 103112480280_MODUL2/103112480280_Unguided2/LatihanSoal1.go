//Nama : Anggun Wahyu W. (103112480280)
package main 

import "fmt"

func main() {
    var tahun int

    fmt.Print("Tahun: ")
    fmt.Scan(&tahun)

    kabisat := false

    if tahun%400 == 0 {
        kabisat = true
    } else if tahun%100 != 0 && tahun%4 == 0 {
        kabisat = true
    }

    fmt.Println("Kabisat:", kabisat)
}