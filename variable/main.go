package main

import (
	"fmt"
)

// Variables are memory addresses with identifiers contain values which can change but must be of the same type. variables are lexically scoped
func main() {
	/*
		Variables declared inside this function are only available within the main function's scope
		variable = [memory address] ---> value

		var identifier type = value
		var identifier = value
		identifier := value
	*/

	// Variable with identifier name 'a' is declared as an int type variable and initialized with its zero value which is '0'
	/*
		int -> 0
		float -> 0.0
		string -> ""
		boolean -> false
		pointer -> nil
		struct -> {0}
	*/
	var a int

	// Variable 'b' declared as string type 'Explicitly assigned' with a string value of 'hello'
	var b string = "hello"

	// Implicitly assigned variable which remains untyped until during compile time when its value will determine its type 'boolean type'
	// automatic type inference
	var c = true

	var (
		// parallel and multiple variable assignment
		d, e, f = false, "yay", 90
	)

	// Initializing declaration of variable identified as h with int type value of 430
	// This only works in a local scope not global or package scope
	h := 430

	fmt.Println(a, b, c, d, e, f, g, h)

	// Reveal the address which contains content of variable identified as 'a'
	fmt.Println(&a)

	// Pointer variable is that which points to a memory address containing a value of specific type
	// var identifier *type = memory address
	var i *int = &a

	fmt.Println(i)
}

// globally scoped or package scoped variable which is available all through out the package
var g = 900
