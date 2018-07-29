package main

import (
	"fmt"
	"math"
)

/* Basic data types
These are the basic types which can be assigned to a variable
numbers = [float, complex and integers]
strings
boolean
*/

func main() {
	// NUMBERS these are types which can hold numeric values such as integers or decimals (floating point values)

	// INTEGERS are whole numbers (positive or negative) where the variable has a defined size of storage (8, 16, 32, 64) bits wide

	// SIGNED INTEGERS can have (-ve or +ve) integer numbers ranging from -2^(n-1) to 2^(n-1) - 1
	var a1 int8 = -128                  // range from -128 to 127
	var a2 int16 = 32767                // range from -32768 to 32767
	var a3 int32 = -2147483648          // range from -2147483648 to 2147483647
	var a4 int64 = 9223372036854775807  // range from -9223372036854775808 to 9223372036854775807
	var a5 int64 = -9223372036854775808 // range from -9223372036854775808 to 9223372036854775807

	// SIGNED NUMBERS are represented in 2's complement form -2^(n-1) to 2^(n-1) -1
	fmt.Println(a1, a2, a3, a4, a5)

	// UNSIGNED INTEGERS these are only positive values that start from 0 to 2^(n) - 1
	var a6 uint8 = 255                   // range from 0 to 127
	var a7 uint16 = 65535                // range from 0 to 32767
	var a8 uint32 = 4294967295           // range from 0 to 2147483647
	var a9 uint64 = 18446744073709551615 // range from 0 to 18446744073709551615

	//By default go compilers will have a variable as type int if not specified

	fmt.Println(a6, a7, a8, a9)

	// Special int type values
	// CHARACTER TYPE
	var a10 rune = 'a' // rune types will store unicode character point value [int32]
	var a11 byte = '*' // same as above [uint8]

	fmt.Println(a10, a11)

	// FLOATING POINT VALUES are values which are represented with a decimal point and can be small or large and are of two types
	// either float32 32 bit wide precision and float64 which is 64 bit wide precision
	// Values can either be positive or negative
	var a12 float32 = math.MaxFloat32
	var a13 float64 = math.MaxFloat64 * -1.0

	//important functions -> float64(value) and float32(value) changes value passed to functions's argument and converts to apporpriate type

	fmt.Println(a12, a13)

	// COMPLEX NUMBERS are values whose components contain a floating point value (real value) and an imaginary component
	// imaginary literal are represented withan i
	var a14 complex64 = 1 + 2i                      // 1 is the real part and 2i is the imaginary part
	var a15 complex128 = complex(a13, float64(a12)) // complex numbers can be created using complex function
	fmt.Println(a14, a15)

	// BOOLEAN VALUES are types which have two states for value which is either true or false
	var a16 bool = true
	var a17 bool = false

	fmt.Println(a16, a17)

	// STRING VALUE are types which contain a sequence of UTF-8 characters with a variable width
	// strings are immutable sequence of chacters which cannot be changed

	var a18 string = "interpreted string here"                    // interpreted string since the characters are surrounded with double quotes ""
	var a19 string = `Raw string
	which can have multiple
	lines` //raw string since characters are surrounded with back ticks ``
	fmt.Println(a18, a19)

	//characters of a string can be accessed using index
	fmt.Println(a18[2])
}
