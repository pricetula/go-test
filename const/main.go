package main

import (
	"fmt"
)

// Constant identifiers contain values which cannot be changed during compile time
func main() {
	// Explicit type assignment of a constant named 'a' as a string type
	const A string = "hello"

	// Implicit type assignment of a constant named 'b' at compile time. constant 'b' remains untypes until compile time
	const B = "world"

	// Constants must be assigned values which are available during compile time
	const C = B

	// Multiple assignments
	const D, E, F = 1, true, "noo"
	const (
		G, H, I = 1, 2, 3
	)

	// Encouraged to use all caps to name constants
	fmt.Println(A, B, C, D, E, F, G, H, I)
}
