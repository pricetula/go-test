package main

import "fmt"

/*
	Slices are a reference to an underlying array in go ie has no set length during definition stage
*/

func main() {
	/*
		Empty int type slice, set without any lenght or elements.
		Memory is allocated
	*/
	slice1 := []int{}
	/*
		slice2 variable declared to be a slice of int type.
		Memory is not allocated this moment
	*/
	var slice2 []int
	/*
		slice3 variable declared to be a slice of int type and initialized with int type values
		Memory is allocated this moment
	*/
	var slice3 []int = []int{1, 2, 3}
	/*
		slice4 variable declared to be of int type having an initial length of 2 and a capacity of 4
		this means the underlying array to which this slice is referencing has a length of 4 elements
		Memory is allocated
	*/
	var slice4 []int = make([]int, 1, 4)
	/*
		slice5 variable declared to be of int type having an initial length of 2 and a capacity of 2
		this means the underlying array to which this slice is referencing has a length of 4 elements
		Memory is allocated
	*/
	var slice5 []int = make([]int, 2)
	slice5[1] = 90
	fmt.Println("slice1 =", slice1)
	fmt.Println("slice2 =", slice2)
	fmt.Println("slice3 =", slice3)
	fmt.Println("slice4 =", slice4)
	fmt.Println("len(slice4) =", len(slice4))
	// Getting capacity of a slice using global func cap
	fmt.Println("cap(slice4) =", cap(slice4))
	fmt.Println("slice5 =", slice5)
	fmt.Println("len(slice5) =", len(slice5))
	fmt.Println("cap(slice5) =", cap(slice5))
	// One can add a new element using global func append(slice, value)
	slice5 = append(slice5, 100)
	fmt.Println("slice5 =", slice5)
	fmt.Println("len(slice5) =", len(slice5))
	fmt.Println("cap(slice5) =", cap(slice5))

	// Spread over the elements of slice4 and appended them with elements from slice5 and initialized to slice6
	var slice6 []int = append(slice5, slice4...)
	fmt.Println("slice6 =", slice6)

	// slice6[0:0] means create a slice with elements from 0 index to 0 - 1 index hence its an empty int slice
	// the below instruction will create a copy of slice6 onto slice7
	var slice7 []int = append(slice6[0:0], slice6...)
	fmt.Println("slice7 =", slice7)
	// Slice starting from index 1
	fmt.Println("slice7[1:] =", slice7[1:])
	// Slice till index "2 - 1" = 1
	fmt.Println("slice7[:2] =", slice7[:2])
	// Slice between index 1 and 2
	fmt.Println("slice7[1:3] =", slice7[1:3])

	// SNew slice without element at index 1
	slice8 := append(slice7[:1], slice7[2:]...)
	fmt.Println("slice8 =", slice8)

	// For loop in range of a slice
	for key, val := range slice8 {
		fmt.Println("Key =>", key, "value=>", val)
	}
}
