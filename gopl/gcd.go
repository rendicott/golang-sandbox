/* GCD function from 
"The Go Programming Language"
section 2.4.1
*/

package main

import (
    "fmt"
)

func gcd(x, y int) int {
    for y != 0 {
        // Adding some debug to show what it's doing
        fmt.Println("before x:",x)
        fmt.Println("before y:",y)
        x, y = y, x%y
        fmt.Println("after x:",x)
        fmt.Println("after y:",y)
        fmt.Println("--------")
    }
    return x
}

func main() {
    fmt.Println("GCD:", gcd(124,10364))
}
