package main

import (
	"fmt"
)

func main() {
	/*
		IF ELSE statement
		=====================
		if condition {
			do something
		} else if another_condition {
			do something else
		} else if (yet_another_condition) {
			This condition van use brackets too
		}  else {
			do this if none pf the above conditons work
		}

		SWITCH statement
		======================
		switch variableName {
		case val1: {
			block to execute
		}
		case val2, val3, val4:
			execute line
		default:
			execute line or block
		}

		FOR loop statement
		======================
		for initialization; condition; modification (mostly increment or decrement) {
			execute block for the loop
		}
		initializations are always in short hand format such as variableName := 0
		modifications only run after the first block execution is finished
	*/

	for a := 0; a < 10; a++ {
		if a < 3 {
			fmt.Println("[ifelse] Less than 3")
		} else if a == 2 {
			fmt.Println("[ifelse] yaay 2")
		} else if a > 5 {
			fmt.Println("[ifelse] More than 5")
		}

		switch a {
		case 1:
			fmt.Println("[switch] Yoo this is 1")
		case 3:
			fmt.Println("[switch] Dayyuuummm 3")
		default:
			fmt.Println("[switch] :(")
		}

		fmt.Println(a)
	}

	// multiple counters on for loop
	for a, b := 0, 20; a != b; a, b = a+1, b-1 {
		fmt.Println("Balancing", a, b)
	}

	// loop through a collection
	str := "Hello world"
	for key, value := range str {
		fmt.Printf("> %d is %c \n", key, value)
	}

	/*
		INFINITE LOOPS
		===============
		for i := 0; ; i ++ {
			infinite loop with counter variable
		}

		for true {
			inifite loop with true boolean set
		}

		for ;; {
			inifite loop with no conditions set
		}

		for {
			infinite loop with nothing
		}
	*/

	a := 0
	for { // Infinite loop
		if a == 21 {
			fmt.Println(a)
			break // Break statement used to break from a for loop
		} else {
			a++
		}
	}

	for { // Infinite for loop
		if a > -40 {
			a-- // reduce value of a
			fmt.Println(a)
			continue // Continue statement used to start program execution to top of block
		}
		break // Break statement will not run if continue statement was executed
	}

	a = 1000
CRAZYLABEL:
	if a > 100 {
		a -= 100
		fmt.Println("Strange loop", a)
		goto CRAZYLABEL // goto label will cause code execution to go back to label
	}
}
