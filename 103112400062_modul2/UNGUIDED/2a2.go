//M.DAVI ILYAS RENALDO
//103112400062

package main

import "fmt"

func main() {
    var tahun int
    fmt.Scan(&tahun)
    
    kabisat := tahun %400 == 0 || (tahun %4 == 0 && tahun%100 !=0)
    fmt.Printf("tahun: %d\nkabisat: %t\n", tahun, kabisat)
}