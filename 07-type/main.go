package main

import (
	"fmt"
	"reflect"
)

// Stucts are copied in memory instead of referenced
type district struct { // Can hold different types of data, anything we want
	name         string `max:"100"` // Setting a tag with space delim, key with :, value in quotation marks
	population   int
	coronaCases  int
	coronaDeaths int
} // A struct can embed other structs
type country struct {
	district // Values are the same as in district
}

func main() {
	Bern := district{name: "Bern", population: 1039101, coronaCases: 1708, coronaDeaths: 55}
	Switzerland := country{district: district{name: "Switzerland", population: 8643685, coronaCases: 29586, coronaDeaths: 1423}}
	fmt.Println(Bern)
	fmt.Println(Switzerland)
	// Reflection is used to get tags and parse them in the code
	// GO does not care about tags
	// This would be used in a validation lib
	t := reflect.TypeOf(district{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}
