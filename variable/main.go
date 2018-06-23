package main

import (
	"fmt"
)

// Variable identifiers contain values which can change but must be of the same type. variables are lexically scoped
func main() {
	// Variables declared inside this function are only available within the main function's scope

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
		d, e, f = false, "yay", 90
	)

	fmt.Println(a, b, c, d, e, f, g)
}

// globally scoped or package scoped variable which is available all through out the package
var g = 900
