/*
In Go, an array is a numbered sequence of elements of a specific
length.
*/

package main

import "fmt"

func main() {
	/*
		Here we create an array a that will hold exactly 5 ints.
		The type of elements and length are both part of the array’s
		type. By default an array is zero-valued, which for ints
		means 0s.
	*/
	var a [5]int
	fmt.Println("emp:", a)

	/*
		We can set a value at an index using the
		array[index] = value syntax, and get a value
		with array[index].
	*/
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	// The builtin len returns the length of an array
	fmt.Println("len:", len(a))

	// use this syntax to declare and init in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	/*
		Array types are one-dimensional, but you can compose types
		to build multi-dimensional data structures
	*/
	var twoD [5][5]int
	fmt.Println("2d: ", twoD)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = i + j
			fmt.Println("2d-working: ", twoD)
		}
	}
	fmt.Println("2d: ", twoD)
	/*
		Note that arrays appear in the form [v1 v2 v3 ...] when
		printed with fmt.Println.
	*/
}

/*
You’ll see slices much more often than arrays in
typical Go. We’ll look at slices next.
*/
