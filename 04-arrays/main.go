package main

import (
	"fmt"
)

func main() {

	// Basic Array declaration
	grades := [3]float32{1.1, 2.2, 3.3}
	fmt.Printf("Grades %v\n", grades)

	// String array example
	students := [7]string{"Alice", "Bob", "Steve", "Peter", "Arnold"}
	fmt.Printf("Students %v\n", students)

	// Indexing
	fmt.Printf("Number of Students %v\n", len(students))
	fmt.Printf("Student #1: %v\n", students[0])

	// Copy and reference arrays
	gradesCopy := grades // Is a full copy of the array, slow because it has to copy all data
	gradesRef := &grades // Is a reference to the array, change values here change origin values too

	fmt.Printf("Grades Copy %v\n", gradesCopy)
	gradesRef[0] = 1.2
	fmt.Printf("Grades referenced %v\n", gradesRef)
	fmt.Printf("Grades old but referenced %v\n", grades)

	// Slices are arrays without predefined sizes
	sliceA := []int{1, 2, 3}
	fmt.Printf("Size of slice A: %v\n", len(sliceA))
	fmt.Printf("Capacity of slice A: %v\n", cap(sliceA))

	// Slices can reference and split arrays
	sliceB := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceC := sliceB[:]   // All elements
	sliceD := sliceC[3:]  // Elements excluding start (element 4)
	sliceE := sliceB[:6]  // Elements including end
	sliceF := sliceB[3:6] // From (excluding) to (including)
	// Updating one value
	sliceB[5] = 0
	// All of the slices are reference to the original array
	fmt.Println(sliceB)
	fmt.Println(sliceC)
	fmt.Println(sliceD)
	fmt.Println(sliceE)
	fmt.Println(sliceF)

	// Define a slice with make instead of definition
	// Can be used to optimize code to assign enough memory for estimated size of array
	sliceG := make([]int, 3, 100)
	fmt.Printf("Length of G: %v\n", len(sliceG))
	fmt.Printf("Capacity of G: %v\n", cap(sliceG))
	sliceG = append(sliceG, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("Length of G: %v\n", len(sliceG))
	fmt.Printf("Capacity of G: %v\n", cap(sliceG))
	fmt.Println(sliceG)

	// Playing with append, slicing and removing elements
	sliceH := []int{1, 2, 3, 4, 5, 6, 7}
	sliceI := sliceH[:len(sliceH)-1] // Removes the last element
	sliceJ := sliceH[1:]             // Removes the first element
	fmt.Println(sliceI)
	fmt.Println(sliceJ)
	// Using append to select the middle part
	sliceK := append(sliceH[:1], sliceH[5:]...)
	fmt.Println(sliceK)
}
