package main

import (
	"fmt"
	"strconv"
)

var d int = 0 // Declare d globally (not only in function main)

// Defining variables the fast way (multiple variables in a block)
var (
	hello string  = "World" // Hello World
	foo   string  = "Bar"   // Foo Bar
	e     float32 = 42.2    // 42 is the answer to the universe
)

// Exported can be used outside of the scope in package level
var Exported int = 0 // Upper case name is exported out of package level

// Multiple blocks of variables can be defined
var (
	counter int = 0
)

func main() {
	// 3 Ways to assign variables, predefined values, undefined and short
	var a int = 0 // Preassign value with defined type
	var b int     // Use when you do not assign value from beginning
	c := 0        // Compiler assumes typing
	// Output defined variables first
	fmt.Println(a, b, c) // Output 0 0 0
	// Show Variable Types
	fmt.Printf("a %v, %T\n", a, a)
	fmt.Printf("b %v, %T\n", b, b)
	fmt.Printf("c %v, %T\n", c, c)
	// Try to reassign C as a float
	c = 0.0
	fmt.Printf("c float %v, %T\n", c, c) // C is still an integer, Go Compiler sets typing on first init
	// Reassigning variables with the same key does not work
	// counter := 11
	// This variable is unused, compiler will throw error: unused declared but not used
	// var unused int

	// Doing a little bit of conversion
	var f float64 = float64(a) // Conversion did not loose data, because we converted int to float
	var g int = int(e)         // Conversion did loose data, because we can not convert float to int (but does not throw exception)
	fmt.Printf("f %v, %T\n", f, f)
	fmt.Printf("g %v, %T\n", g, g)
	// Strings
	var h string = string(g) // Trying to convert an int directly to a string
	// Does not work because it converts int to the corresponding char
	fmt.Printf("h %v, %T\n", h, h)
	// Instead we have to use strconv package
	h = strconv.Itoa(g) // Correct way to convert int to string
	fmt.Printf("h %v, %T\n", h, h)

	// Implicit conversion does not work
	// You have to tell the compiler to convert explicit
	// ! g = e
	// g = int(e)

	// Global variables can be shadowed in functions

	// All variables must be used

	// 3 Scopes; Public (Uppercase), Package (Lowercase) and Block (Either)
	// Naming conventions, Pascal and camelCase

	// Boolean
	var i bool = false // Default value is false (uninitialized)
	// i = true
	fmt.Printf("i %v, %T\n", i, i)
	j := i == true // Tests this value and sets j accordingly as a boolean
	k := i == false
	fmt.Printf("j %v, %T\n", j, j)
	fmt.Printf("k %v, %T\n", k, k)

	// Numeric Types
	// Signed
	var l int = 42   // Int is undefinded size, could be any size
	var m int8 = 8   // 8 Bit integer
	var n int16 = 16 // 16 Bit integer
	var o int32 = 32 // 32 Bit integer
	var p int64 = 64 // 64 Bit integer
	// Unsigned
	var q uint = 42
	var r uint32 = 32
	var s byte = 8

	// Small calculations
	// Direct conversions do not work
	// ! fmt.Println(l - m)
	fmt.Println(l + int(m))
	fmt.Println(n - int16(o))
	fmt.Println(p * int64(q))
	fmt.Println(r / uint32(s))

	t := 8 // 0100
	u := 4 // 0010
	// Boolean operations
	fmt.Println(t & u)  // 0000
	fmt.Println(t | u)  // 0110
	fmt.Println(t ^ u)  // 0110
	fmt.Println(t &^ u) // 1001
	fmt.Println(t << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(t >> 3) // 2^3 / 2^3 = 2^0 = 1

	e = 3.14E10 // Can use E notation
	fmt.Println(e)

	// Complex numbers
	var v complex64 = 2i
	var x complex64 = 3 + 12i
	fmt.Println(x + v)

	// Text (UTF-8)
	h = "This is a string"
	fmt.Println(h)
	// A string is an array
	fmt.Println("Letter 3 of the String is " + string(h[2])) // Because h2 is a byte (8) we have to convert it back to string

}
