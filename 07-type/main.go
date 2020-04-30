package main

import "fmt"

type canton struct { // Can hold different types of data, anything we want
	name         string
	population   int
	coronaCases  int
	coronaDeaths int
} // A struct can embed other structs

func main() {
	Bern := canton{name: "Bern", population: 1039101, coronaCases: 1708, coronaDeaths: 55}
	fmt.Println(Bern)
}
