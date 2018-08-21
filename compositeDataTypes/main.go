package main

import (
	"fmt"
)

/*
COMPOSITE DATA TYPES
-----------------------
These are data types which are created by combining basic data types to create a more complex type
Arrays, Structs, Slices and Maps
*/

func main() {
	/* ARRAY DATA TYPE created by combining multiple elements of the same basic type whose length is fixed */

	var a1 [4]int // Declared a variable called a1 as an array of int type having a fixed length of 4. The array only holds int type elements
	a1[2] = 90
	fmt.Println(
		"a1",
		a1,
		a1[2],   // Get an element value using index
		len(a1), // Use len function to determine length of array (the total number of elements present in the array)
		cap(a1), // determine capacity of array. (total elements which the array can hold)
	)

	var a2 [2]string = [2]string{"Hello", "World"} // Declaring variable a2 as a string type array having a length 2 is initialized with values
	fmt.Println("a2", a2)

	var a3 = [2]int{16} // Variable a3 will be declared as an array of int type having a length 2 but initialized with an element 1 (Infered)
	fmt.Println("a3", a3)

	// ... ellipsis is used to denote that the array length will be how many elements it will be initialized with
	var a4 = [...]bool{true, false, true} // Variable a4 is an array of length 3.
	fmt.Println("a4", len(a4))

	var a5 = [...]string{2: "Hello", 4: "World"} // Initialized array with string values at specified indicies. The length will be determined by the highest index during initialization
	fmt.Println("a5", a5, len(a5))

	// comparing arrays
	a6 := [...]int{1, 2}
	a61 := [2]int{1: 2, 0: 1}
	fmt.Println(a6 == a61)
}
