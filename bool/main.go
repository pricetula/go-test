package main

import (
	"fmt"
)

// Boolean values are those having two states either true or false
func main() {
	a, b := false, false

	// reassigning variable a with the opposite boolean value of variable b
	// not operator !booleanIdentifier
	a = !b

	fmt.Println(
		a,
		b,

		// Comparison operators
		// Equality operator
		a == b,
		// Not equal operator
		a != b,

		// Logical operators
		// not operator
		!a,
		// and operator
		a && b,
		// Or operator
		a || b,
	)
}
