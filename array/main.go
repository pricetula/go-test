package main

import "fmt"

func main() {
	// variable definied to be of int type array having a length of 0, no item can be added into the array since its zero length
	var arr1 = [0]int{}
	// arr1[0] = 9 throws error
	/*
		arr1 variable is initialized with an array having 3 int type items
		'...' elipsis is used when one doesn't want to specify the length of the list it will be obtained from values within '{}'
	*/
	var arr2 = [...]int{1, 2, 3}
	/*
		definied variable of string of array type with length of 2 is initialized with two string elements
	*/
	var arr3 [2]string = [2]string{"", ""}
	arr3[1] = "mango"
	arr3[0] = "hello"
	/*
		array of int type having a length of 2 is initialized with its zero values "", 0, false
	*/
	var arr4 [2]int
	/*
		Multi dimensional arrays have [size1][size2]...[sizeN]Type{
			[size1]Type { val1, ... valN }
			[size2]Type { val1, ... valN }
			...
			[sizeN]Type { val1, ... valN }
		}
	*/
	arr5 := [2][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}

	fmt.Println("arr1 = ", arr1)
	fmt.Println("arr2 = ", arr2)
	fmt.Println("arr3 = ", arr3)
	fmt.Println("arr4 = ", arr4)
	fmt.Println("arr5 = ", arr5)
	fmt.Println("len(arr2) ", len(arr2))
}
