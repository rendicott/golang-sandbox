package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [5][5]int
	fmt.Println("2d: ", twoD)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = i + j
			fmt.Println("2d-working: ", twoD)
		}
	}
	fmt.Println("2d: ", twoD)
}
