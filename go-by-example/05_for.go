/*
for is Go’s only looping construct. Here are three basic
types of for loops.
*/

package main

import "fmt"

func main() {
	// the most basic type with a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("----------")
	// a clasic iniitial/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("----------")
	/*
		for without a condition will loop repeatedly until you
		break out of the loop or return from the enclosing
		function
	*/
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("----------")

	// you can also continue to the next iteration of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

/*
We’ll see some other for forms later when we look at range
statements, channels, and other data structures.
*/
