# 05-Sudoku

## Goal of this session

I am hacking a tool that solves Sudokus in Go. This was done in one Go (~5 Hours).

## How does it work?

1. Reading in a human readable source file (.sudoku) as source
2. Running a reverse backtrack algorithm to find a possible solution
3. Outputing the solution as a console output in the .sudoku format to stdout

## Running the example

`go run main.go {filename}`
Example is provided in example.sudoku

You can create your own sudokus to get solved in the following format:

```
0 0 8 3 1 2 0 9 7
0 4 0 0 5 0 0 0 6
0 9 1 0 7 4 0 8 3
0 0 0 0 6 0 0 4 5
0 0 4 1 0 5 0 0 0
0 8 0 2 0 7 0 0 0
0 0 0 0 0 0 0 0 9
8 0 0 4 0 0 7 0 0
9 6 0 0 0 0 0 0 0
```

## How can we extend it

1. Creating our own sudokus

## Result

I had a lot of fun, creating my own sudoku solver as a small project to learn go. It was quite challenging to understand backtracking.
The funniest error that occured, was that I could not use a pointer as a function parameter. I had to use * in front of the type to ensure the compiler accepted it.

### Lessions learned

* A lot about typing in Go
* A ton about returning, defining functions and structure in go applications
* A lot of funny errors and first usage of file io in go
* I see now that Go does not handle a lot itself in runtime, possible errors have to be handled by the code itself
