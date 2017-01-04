package main

import "fmt"

func f(c *int) int {
	var b int = *c
	*c += 5
	return b
}

func fc(a int) int {
    return a += 20
}


func main() {
	var p int = 5
	fmt.Println("var p:", p)

	r := &p
	fmt.Println("var r:", r)
	fmt.Println("var *r:", *r)

	fmt.Println("Before f(), p was:", p)
	fmt.Println("running f()..., return is:", f(&p))
	fmt.Println("After f(), p is:", p)
	fmt.Println("running f()..., return is:", f(&p))
	fmt.Println("After f(), p is:", p)

    var j int = 0
    fmt.Println("run1 fc(j) =", fc(j))
    fmt.Println("run2 fc(j) =", fc(j))
}
