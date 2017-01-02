// Go supports methods defined on struct types.

package main

import "fmt"

type rect struct {
	width, height int
}

// This area method has a receiver of type *rect
func (r *rect) area() int {
	return r.width * r.height
}

/*
Methods can be defined for either pointer or value receiver
types. Here's an example of a value receiver.
*/
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	rr := rect{width: 10, height: 5}
	// here we call the two methods defined for our struct
	fmt.Println("area: ", rr.area())
	fmt.Println("perim: ", rr.perim())

	/*
		Go automatically handles conversion between values and
		pointers for method calls. You may want to use a pointer
		receiver type to avoid copying on method calls or to
		allow the method to mutate the receiving struct.
	*/
	rp := &rr
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}

/*
Next we’ll look at Go’s mechanism for grouping and naming
related sets of methods: interfaces.
*/
