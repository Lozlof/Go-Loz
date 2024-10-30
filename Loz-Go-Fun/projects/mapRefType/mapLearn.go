package main

import (
	"fmt"
)

// In Go, a map is a built-in data structure that stores key-value pairs. 
// It’s similar to a dictionary in other languages. 
// Maps are useful when you need to associate a set of unique keys with values, where each key is mapped to a specific value.

func main() {
	// In Go, declaring and initializing are two different steps when working with variables, maps, arrays, or other data structures.

	// Declaration: 
	// Means specifying a variable’s name and type, but it doesn’t assign it a value (except possibly the zero value).
	// Declaration tells the compiler that a variable exists, but it doesn’t allocate memory for a value yet, except for zero values in Go.
	
	// Initializing:
	// Initialization means assigning an initial value to a variable. 
	// This can be done at the time of declaration or afterward.
	// Initializing a variable gives it a specific value and allocates the memory needed for that value.

	fmt.Println("Declaring theMap")
	// Declaring a map.
	var theMap map[string]int
	// The KEY type is string.
	// The VALUE is INT
	fmt.Println("theMap is length:", len(theMap))

	fmt.Println("Declaring and Initializing theMapTwo")
	// Declare and initialize an empty map.
	var theMapTwo = make(map[string]int)
	// The make built-in function allocates and initializes an object of type slice, map, or chan (only). 
	// Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. 
	// The specification of the result depends on the type
	// More info on make: https://pkg.go.dev/builtin
	fmt.Println("theMapTwo is length:", len(theMapTwo))

	fmt.Println("Declaring and Initializing theMapThree with values")
	// Declare and initialize with values
	theMapThree := make(map[string]int)
	theMapThree["shoes"] = 13
	theMapThree["socks"] = 45
	theMapThree["shirts"] = 3
	theMapThree["sweater"] = 13

	// If you know the key, you can access and print the value directly.
	fmt.Println("theMapThree values:")
	fmt.Println("shoes:", theMapThree["shoes"])
	fmt.Println("socks:", theMapThree["socks"])
	fmt.Println("shirts:", theMapThree["shirts"])
	fmt.Println("sweater:", theMapThree["sweater"])
}