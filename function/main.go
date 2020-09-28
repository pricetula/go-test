package main

import (
	"fmt"
)

/*
 FUNCTIONS
 ==========
 Are blocks for executing code encapsulated within {}
 when functions are called or executed they will finish once last statement has been executed or a value is returned
 functions can have arguments for passing values to the inside block
 Once it has finished execution the code instruction execution continues outside its block

 func functionName (argument1, argument2, ..., argumentN) {
	 code block to be executed
 }

 functionName() called function is executed
*/
func main() {
	a1()
	a2(2)
	a3("hello", "world")
	a4(400, "hello world")
	a5(1, 2, 3, 4)
	a6(1, true, "Hello")
	fmt.Println(a7())
	fmt.Println(a8())
	fmt.Println(a9())
	fmt.Println(a10())
	fmt.Println(a11())

	count := 0
	a12(count)

	a13()()

	fmt.Println(a14()(3))

	replicateInstance := replica(1)  // value 1 initialized
	fmt.Println(replicateInstance()) // returned function will multiply what was insitialized with before
	fmt.Println(replicateInstance())
	fmt.Println(replicateInstance())
	fmt.Println(replicateInstance())
	fmt.Println(replicateInstance())

	a15()

	a17()
}

func a1() {
	fmt.Println("a1")
}

func a2(v int) { // Argument of type int required
	fmt.Println("a2 value >", v)
}

func a3(v, x string) { // Two arguments of type string required
	fmt.Println("a3 values >", v, x)
}

func a4(v int, x string) { // Two arguments of type int and string required
	fmt.Println("a4 values >", v, x)
}

func a5(params ...int) { // variadic function that accepts unlimited number of int type arguments
	fmt.Println(params) // The arguments are bundled as an array of int type values [len(params)]int
}

func a6(params ...interface{}) { // Variadic function that accepts unlimited number of empty interface type (not typed)
	fmt.Println(params)
}

func a7() string { // Function that returns a single string type value
	return "Hello World"
}

func a8() (string, int) { // Function that returns multiple values of string and int types and must be returned in stated position
	return "Hello World", 4000
}

func a9() (a int) { // named return type function. A variable is defined and initialized as we declare it to be the value to be returned
	a = 20 // variable a is already defined as the expected value to be returned

	return // we dont specify what to return since we already declared variable a of type int should be returned
}

func a10() (a int, b string) { // multiple named return values
	a = 90
	b = "Hello worlD"

	return
}

func a11() (a, b string) { // multiple named return values both of type string
	a = "Herro"
	b = "Goobye"

	return
}

func a12(c int) { // Recursive function calls it self
	if c < 2 {
		a12(c + 1) // recursive call to itself. The last called function will resolve first
	}
	fmt.Println("Recurs::", c)
}

func a13() func() { // function that returns an anonymous function
	return func() {
		fmt.Println("Returned Function")
	}
}

func a14() func(int) int { // function that returns an anonymous function which has an int type argument and returns an int type value
	return func(a int) int { // anonymous function returned
		return 5 * a
	}
}

func replica(a int) func() int {
	return func() int {
		a *= 2

		return a
	}
}

func a15() { // defer is a label which means a block of code is run right before the called function exits
	defer func() { // function is defered hence will run last
		fmt.Println("I will actually run last")
	}()

	fmt.Println("I will run first I promise dude")
}

func a16() {
	panic("Hey i paniced")
}

func a17() {
	defer func() {
		// Defered function runs at the end of this function block and will get recovored panic
		panicVal := recover()
		if panicVal != nil {
			fmt.Println("Recovered Panic message:: ", panicVal)
		}
	}()
	// This function will panic
	a16()
}
