package main

import "fmt"

func main() {

	cantonPopulations := map[string]int{
		"Zurich":                 1538848,
		"Bern":                   1039101,
		"Waadt":                  804861,
		"Aargau":                 685642,
		"St.Gallen":              510670,
		"Genf":                   504031,
		"Luzern":                 413048,
		"Tessin":                 351471,
		"Wallis":                 345394,
		"Fribourg":               321717,
		"Basel-Landschaft":       289404,
		"Thurgau":                279493,
		"Solothurn":              275177,
		"Graubuenden":            198988,
		"Basel-Stadt":            195783,
		"Neuenburg":              176467,
		"Schwyz":                 160457,
		"Zug":                    127612,
		"Schaffhausen":           82337,
		"Jura":                   73563,
		"Appenzell Ausserrhoden": 55432,
		"Nidwalden":              43076,
		"Glarus":                 40852,
		"Obwalden":               37924,
		"Uri":                    36694,
		"Appenzell Innerhoden":   16127,
	}

	// Easy conditionals
	if true {
		fmt.Println("This is always true")
	}
	if false {
		fmt.Println("This will never been shown")
	}

	// Run an initializer to generate a result and run the if statement, when a valid value has been generated
	if canton, ok := cantonPopulations["Bern"]; ok {
		fmt.Println(canton)
		if canton < 1000000 {
			println("It is smaller than 1 Million")
		}
		if canton > 1000000 {
			println("It is bigger than 1 Million")
		}
		if canton == 1000000 {
			println("It is exactly 1 Million")
		}
	} // Canton is only defined in the if statement

	// Logical operators
	if true && false {
		// Will not execute, since not both are true
		fmt.Println("This will never been shown")
	}
	if true || false {
		fmt.Println("This will always be shown, because OR allows only one to be true")
	}
	if true && !false {
		fmt.Println("This will always be shown, because we inverted false with NOT")
	}

	// Easy switch statement example
	switch i := 1 + 2; i { // Use initializer
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3, 4:
		fmt.Println("Three or Four")
	default:
		fmt.Println("Other Number")
	}
	// Tagless switch statement example
	i := 10
	switch {
	case i > 10:
		fmt.Println("I is bigger than 10")
	case i < 10:
		fmt.Println("I is smaller than 10")
	default:
		fmt.Println("I is 10")
	}
	// Type switch statement example
	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("J is an integer")
	default:
		fmt.Println("J is not an integer")
		break
		// fmt.Println("This will not be shown")
		// Code is unreachable
	}
}
