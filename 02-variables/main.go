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

}
