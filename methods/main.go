package main

import "fmt"

func main() {
	c := Calc{
		a: 9,
		b: 91,
	}
	c.Add()
	fmt.Println(c.result)
	c.Subtract()
	fmt.Println(c.result)
	c.Multiply()
	fmt.Println(c.result)
	c.Divide()
	fmt.Println(c.result)
	c.MaddAdd()
	fmt.Println(c.result)
}
