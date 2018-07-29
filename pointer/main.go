package main

import (
	"fmt"
)

// Pointer is a variable that contains memory address which points to a value of specific type
// var identifier *pointer-type = [memory address]
// &identifierVariable => [get memory address]
func main() {
	var a *int = &testA

	// variable a has memory address which points to variable testA contents
	fmt.Println(a)

	// variable b is now a pointer since variable a's contents have been copied to b
	// any change to testA variable will also affect that pointed with a and b since they both point to same memory location
	b := a

	fmt.Println(*a, *b, &b) // NB &b is the address of pointer which contains memory address of variable testA

	testA = 239

	// will print 239 since the value pointed by a and b (pointers) has changed to 239
	fmt.Println(*a, *b, &b)
}

// This variable is stored in heap
var testA = 90
