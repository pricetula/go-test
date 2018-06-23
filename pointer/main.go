package main

import (
	"fmt"
)

// Pointer is a variable that contains memory address which points to a value of specific type
// var identifier *pointer-type = [memory address]
// &identifierVariable = [memory address]
func main() {
	var a *int = &testA

	// variable a has memory address which points to variable testA contents
	fmt.Println(a)

	testA = 90

	// variable b is now a pointer since variable a's contents have been copied to b
	// any change to testA variable will also affect that pointed with a and b since they both point to same memory location
	b := a

	fmt.Println(*a, *b, &b) // NB &b is the address of pointer which contains memory address of variable testA
}

// This variable is stored in heap
var testA = 90
