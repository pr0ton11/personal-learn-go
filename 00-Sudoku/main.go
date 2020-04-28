package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var gridBuffer [9][9]int // Holds the current grid, passed to the application (default 0)

func main() {

	// Parsing command line and getting the filename
	filename, err := parseArguments(os.Args[1:])
	if err == nil {
		// Parsing the file given as command line argument
		parseFile(filename)
	} else {
		log.Printf("Warning: Usage go run main.go {filename}")
	}
	// Show the input
	println("INPUT")
	println("-----")
	outputGrid(gridBuffer)
	println("-----")

	// Solving the grid
	println("Solving Sudoku...")
	if solveGrid(&gridBuffer) {
		println("Sudoku was solved sucessfully:")
		outputGrid(gridBuffer)
	} else {
		println("Sudoku could not be solved")
		println("Try another one :-(")
	}
}

// Parses the arguments of the command line
// returns a filename and error
func parseArguments(args []string) (filename string, err error) {
	if len(args) == 0 {
		err = errors.New("Missing argument: {Filename}")
	} else {
		filename = args[0]
	}
	return
}

// Function to check a file, read and parse it to grid
func parseFile(filename string) {
	// Check if file exists
	// Check read permissions
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666) // Open file in read only mode
	if err != nil {                                       // Check if we got an error
		if os.IsNotExist(err) { // If the error was file not existent
			log.Fatal("File " + filename + " does not exist") // Log a fatal error
		}
		if os.IsPermission(err) { // If the error was file not readable
			log.Fatal("File " + filename + " is not readable (permissions)") // Log a fatal error
		}
	}
	// Creating the scanner
	scanner := bufio.NewScanner(file) // Create a new scanner from file
	scanner.Split(bufio.ScanLines)    // Use scanlines (default to split each line)
	// Parsing the file
	for l := 0; l < 9; l++ { // 9 lines to read
		// Running the scanner
		succ := scanner.Scan()
		// Check if the scan was sucessful
		if succ == false {
			// Checking error
			err = scanner.Err()
			if err != nil {
				log.Fatal(err)
			}
		}
		line := scanner.Text() // Read the line as text
		// Use fields function as it returns a split of white spaces from line
		numbers := strings.Fields(line)
		if len(numbers) != 9 { // Sanity check if line does not contain 9 numbers
			log.Fatal("Conversion error on file " + filename + ": Line " + strconv.Itoa(l) + " does not contain 9 numbers")
		}
		for col := 0; col < 9; col++ {
			// l equals line in grid
			// col equals column in grid
			gridBuffer[l][col], err = strconv.Atoi(numbers[col])
			if err != nil { // Checking for conversion errors
				log.Fatal("Conversion error on file " + filename + ": Line " + strconv.Itoa(l) + " does not contain valid integers")
			}
		}
	}
}

// Outputs a grid to stdout
// @param grid, 2d int array representing our custom sudoku type
func outputGrid(grid [9][9]int) {
	for row := 0; row < 9; row++ { // Iterates rows
		for column := 0; column < 9; column++ { // Iterates colums
			fmt.Printf("%d ", grid[row][column]) // Prints the current values
		}
		fmt.Print("\n") // Adds a new line for each row
	}
}

// Checks if the grid is valid
// @param grid, 2d int array representing our custom sudoku type
func validateGrid(grid *[9][9]int) bool {
	// Checking the row
	for row := 0; row < 9; row++ {
		seen := [10]int{} // Create an array with index numbers that we have seen
		for col := 0; col < 9; col++ {
			seen[grid[row][col]]++ // Assign counts of index numbers we have seen
		}
		if containsDuplicates(seen) { // Grid is invalid when it contains duplicates
			return false
		}
	}
	// Checking the column
	for row := 0; row < 9; row++ {
		seen := [10]int{} // Create an array with index numbers that we have seen
		for col := 0; col < 9; col++ {
			seen[grid[col][row]]++ // Assign counts of index numbers we have seen
		}
		if containsDuplicates(seen) { // Grid is invalid when it contains duplicates
			return false
		}
	}
	// Checking the sections
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			seen := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					seen[grid[row][col]]++
				}
				if containsDuplicates(seen) {
					return false
				}
			}
		}
	}
	return true
}

// Checks if a int array contains duplicates
// @param input int array
func containsDuplicates(input [10]int) bool {
	for i, count := range input { // Set index and count for the range in input
		if i == 0 { // Saw index 0 times
			continue
		}
		if count > 1 { // Saw index more than 1 times
			return true
		}
	} // Saw index exactly once
	return false
}

// Checks if we have empty cells
func gridHasEmptyCells(grid *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

// Searches a solution for the grid
func solveGrid(grid *[9][9]int) bool {
	if !gridHasEmptyCells(grid) {
		return true
	}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				for potential := 9; potential >= 1; potential-- {
					grid[row][col] = potential
					if validateGrid(grid) {
						if solveGrid(grid) {
							return true
						}
						grid[row][col] = 0
					} else {
						grid[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return false
}
