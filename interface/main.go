package main

import "fmt"

type moreInfo interface {
	name() string
	height() int
	size() int
}

type heightIncreaseable interface {
	increaseHeight()
}

type nameChangeable interface {
	changeName()
}

type workLifeInfo interface {
	employer() employerDetails
}

type employerDetails struct {
	name     string
	position string
}

type human struct {
	firstName  string
	lastName   string
	lastHeight int
	heapSize   int
}

type snake struct {
	speciesName   string
	headHeight    int
	slitherLength int
}

// Type human implemnting the workLifeInfo interface
// employer implementation
func (h human) employer() (v employerDetails) {
	v = employerDetails{
		name:     "Pack man " + h.firstName,
		position: "Junior",
	}
	return
}

// Type human implementing the moreInfo interface
// name implementation
func (h human) name() string {
	return h.firstName + " " + h.lastName
}

// height implementation
func (h human) height() int {
	return h.lastHeight
}

// size implementation
func (h human) size() int {
	return h.lastHeight * h.heapSize
}

// Type snake implementing the moreInfo interface
// name implementation
func (s snake) name() string {
	return s.speciesName
}

// height implementation
func (s snake) height() int {
	return s.headHeight
}

// size implementation
func (s snake) size() int {
	// reference value s is passed by value
	return s.headHeight + s.slitherLength
}

func (h *human) changeName() {
	h.firstName = "new"
	h.lastName = "name"
}

func (h *human) increaseHeight() {
	if h.lastHeight < 150 {
		h.lastHeight += 10
		fmt.Println("Increasing height of ", h.firstName+" "+h.lastName, " to ", h.lastHeight)
		h.increaseHeight()
	}
}

func printDetails(info moreInfo) {
	fmt.Println("name: ", info.name())
	fmt.Println("height: ", info.height())
	fmt.Println("height: ", info.size())
}

func printEmployerInfo(w workLifeInfo) {
	fmt.Println(w.employer())
}

func changeNameFor(u nameChangeable) {
	u.changeName()
}

func increaseHeightFor(h heightIncreaseable) {
	h.increaseHeight()
}

func main() {
	jane := human{
		"Jane",
		"Wairimu",
		90,
		900,
	}
	python := snake{
		speciesName:   "black python",
		headHeight:    9,
		slitherLength: 90,
	}
	// the jane object must implement all methods specified by moreInfo interface
	printDetails(jane)
	// the snake object must implement all methods specified by moreInfo interface
	printDetails(python)
	// the jane object must implement all methods specified by workLifeInfo interface
	printEmployerInfo(jane)
	increaseHeightFor(&jane)
	changeNameFor(&jane)
	fmt.Println(">>", jane)
}
