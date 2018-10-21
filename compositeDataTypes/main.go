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
	a6 := [...]int{1, 2, 4}
	a61 := [3]int{1: 2, 0: 1, 2: 4}
	fmt.Println(a6 == a61)

	/* SLICE DATA TYPE is a refrence to an underlying array */
	a7 := []int{1, 2} // a slice having capacity of 2 and length of 2. This is created without specifying the length.
	fmt.Println(a7, len(a7), cap(a7))

	var a8 []int = a61[1:]
	fmt.Println(a8, len(a8), cap(a8))

	a81 := a61[1:] // slice initialized with contents of array a61 from index above 1 to the last content
	fmt.Println(a81, len(a81), cap(a81))

	a82 := a61[1:2]
	fmt.Println(a82, len(a82), cap(a82))

	a61[1] = 90 // Any change to the refrenced array will change the contents of slice
	fmt.Println(a81, a82)

	a81[0] = 100 // Any change of the slice content will change the refrence array and any other slice refrencing the same array
	fmt.Println(a61, a81, a82)

	a9 := make([]int, 0, 5)
	a9 = a81
	fmt.Println(a9, len(a9), cap(a9))

	// a91 := make([]int, 0, 5)
	// a91 = a6
	// fmt.Println(a91, len(a91), cap(a91))

	// new(T) allocates memory and initializes value to a zerod type T value and returns a pointer to it
	// make(T) allocates memory and initializes value and returns the initialized value of type T

	// MAP type is a collection of key value pair items
	// var variableName map[keyType]valueType

	a10 := make(map[int]string)

	a10[1] = "hello"
	a10[2] = "world"
	fmt.Println(a10)

	// use the map literal to create a map type value
	a11 := map[int]string{
		1: "ola",
		2: "espanol",
	}
	fmt.Println(a11, a11[2])

	// delete function used to remove an element (key & value pair) from a map
	delete(a11, 1)
	fmt.Println(a11)

	// deleting an element whose key doesn't exist in the map doesn't throw an exception
	delete(a11, 10)
	fmt.Println(a11)

	// when a non existing element is refrenced, its value is takes the 'zeroed' version of type defined :below element with key 9 has an empty string value
	a11[9] = a11[9]
	fmt.Println(a11)

	// Concatenation of initialized value which is an empty string value with 'Hello world' string value
	a11[90] += "Hello world"
	fmt.Println(a11)

	// Iterate through a map using for loop
	for key, value := range a11 {
		fmt.Println(key, value)
	}
}
