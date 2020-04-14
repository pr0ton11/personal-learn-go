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
	// Error would be the default value, because Go assigns 0 to undefined int (nil)
	const (
		err = iota
		male
		female
	)
	// Define a valid gender
	var gender int = male
	// Define an invalid gender
	var invalidGender int
	// Check if the results are what was expected
	fmt.Println(gender == male)
	fmt.Println(invalidGender == female || invalidGender == male)

	// Constants shift index by 5
	const (
		_   = iota + 5
		foo // 6
		bar // 7
	)

	// Example of shifting in constants
	// This works as long as the value is clearly defined on build time
	const (
		_  = iota             // Use default IOTA index to ignore fist item
		KB = 1 << (10 * iota) // Bit shift 1 with 10 * iota
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	// Example to use it as a bitmask for roles
	// Encoded all roles as one single byte instead
	const (
		isAdmin = 1 << iota
		isHQ
		canSeeFinances
		canSeeEurope
		canSeeAsia
		canSeeAmerika
	)
	// Set role and output
	roles := isAdmin | canSeeFinances | canSeeEurope | canSeeAsia
	fmt.Printf("%b\n", roles)

	// Check role
	fmt.Printf("Has Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Can see finances? %v\n", canSeeFinances&roles == canSeeFinances)

}
