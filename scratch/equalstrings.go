package main

import "fmt"
import "strings"

func main() {
    mystring := 
`one
two
three
    four`
    fmt.Println(mystring)

    var slic []string
    slic = append(slic, fmt.Sprintf("one\n"))
    slic = append(slic, fmt.Sprintf("two\n"))
    slic = append(slic, fmt.Sprintf("three\n"))
    slic = append(slic, fmt.Sprintf("    four"))
    var s string
    s = strings.Join(slic, "")
    fmt.Println("---------------------------")
    fmt.Println("'" + s + "'")
    fmt.Printf("mystring = s ? %t\n", mystring==s)
    fmt.Printf("isEqualString(mystring, s) = ?: %t\n", isEqualString(mystring, s))
}

func isEqualString(a, b string) bool {
    r := true
    var lettersA []rune
    var lettersB []rune
    // loop through each char in each word and make a runemap
    for i, r := range a {
        lettersA= append(lettersA, r)
        fmt.Printf("lettersA CHAR %d: %b = %q\n", i, r, lettersA)
    }
    fmt.Println("---------------------------------")
    for i, r := range b {
        lettersB = append(lettersB, r)
        fmt.Printf("lettersB CHAR %d: %b = %q\n", i, r, lettersB)
    }
    
    for i, v := range lettersA {
        if lettersB[i] != v {
            r = false
            fmt.Printf("CHAR %d: %b != %b\n", i, lettersB[i], v)
            break
        }
    }
    return r
}