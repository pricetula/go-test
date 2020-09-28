package main

// Calc structure of an example method thingy
type Calc struct {
	a      int
	b      int
	result int
}

// All methods are best used if the receiver is given a pointer of the type its working with

// Add method allows addition
func (c *Calc) Add() {
	c.result = c.a + c.b
}

// Subtract method allows subtraction
func (c *Calc) Subtract() {
	c.result = c.a - c.b
}

// Multiply method allows multiplication to occur
func (c *Calc) Multiply() {
	c.result = c.a * c.b
}

// Divide function allows one to divide
func (c *Calc) Divide() {
	c.result = c.a / c.b
}

// MaddAdd crazy ting fam
func (c *Calc) MaddAdd() {
	c.Multiply()
	c.a = c.result
	c.Subtract()
	c.b = c.result
	c.Add()
}

func intfy(v float64) int {
	return int(v)
}
