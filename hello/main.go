// specify that this file is the main package "The entry point of go program execution"
package main

// import IO formating package
import "fmt"

// initializing function that runs before main function
func init() {
	fmt.Println("Initializing...")
}

// main function which is the last function that is executed after all other functions and packages imported have been evaluated
func main() {
	fmt.Println("Hello World")
}
