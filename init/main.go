package main

import (
	"fmt"
)

var a string

// Init function is that which is called before the main function and is single threaded.
// mostly used for initialization purposes before main is called
func init() {
	a = "Yooo"
}

func main() {
	fmt.Println(a)
}
