// A goroutine is a lightweight thread of execution

package main

import "fmt"

func f(from string) {
	for i := 0; i < 8; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	/*
	   Suppose we have a function call f(s). Here's how we'd
	   call that in the usual way, running it synchronously.
	*/
	f("direct")

	/*
	   To invoke this function in a goroutine, use go f(s).
	   This new goroutine will execute concurrently with the
	   calling one.
	*/
	go f("goroutine")

	/*
	   You can also start a goroutine for an anonymous
	   function call.
	*/
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	/*
	   Our two function calls are running asynchronously
	   in separate goroutines now, so execution falls through
	   to here. This Scanln code requires we press a key
	   before the program exits.
	*/
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}

/*
When we run this program, we see the output of the blocking
call first, then the interleaved output of the two
goroutines. This interleaving reflects the goroutines being
run concurrently by the Go runtime.

c:\source\go> go run 22_goroutines.go
direct : 0
direct : 1
direct : 2
direct : 3
direct : 4
direct : 5
direct : 6
direct : 7
goroutine : 0
goroutine : 1
goroutine : 2
goroutine : 3
goroutine : 4
goroutine : 5
goroutine : 6
going
goroutine : 7

done
*/
