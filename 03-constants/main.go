package main

import (
	"fmt"
)

func main() {
	const a int = 42 // Constants should still start lower case because we do not want to export it by default
	// Constants usage will be replaced by the compiler at compile time
	// fmt.Println(a) = fmt.Println(42)
	fmt.Println(a)
	// const b float32 = math.Sin(1.57) // Constants can only be declared when the result can be assigned at compile time

	// Special Constant IOTA
	// Works as an Increment for each time IOTA is used
	const (
		b = iota // 0
		c = iota // 1
		d = iota // 2
	)
	fmt.Println(b, c, d)
	// This can be used as a single declaration
	const (
		e = iota // 0
		f        // Will be assigned IOTA by the compiler, 1
	)

	// With this we can create Enumerations
	const (
		err = iota
		male
		female
	)

	var gender int = male
	var invalidGender int = err
	fmt.Println(gender == male)
	fmt.Println(invalidGender == female)

}
