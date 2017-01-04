package main

import "fmt"
import "os"

func main() {
	fmt.Println("hello")
	f, err := os.Open("sample.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(f)
}
