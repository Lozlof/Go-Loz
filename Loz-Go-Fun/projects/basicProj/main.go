// go env GOPATH : View environmental variable GOPATH.
// go env : View all environmental variables.

// go mod init example.com/Loz-Go-Fun : Creates a new module, initializing the go.mod file that describes it.
// go get rsc.io/quote : Changes the required version of a dependency (or adds a new dependency).
// go mod tidy : Removes unused dependencies.
// go build, go test, and other package-building commands add new dependencies to go.mod as needed.
// go run . : Compiles and the Go program in one step.

// Package declaration.
// The main package is special in Go because it's used for programs that are meant to be executable.
// Every Go program that you want to run must have a main package, and within that package, it must have a main function. 
// This is the entry point of your program, and it tells Go where to start executing code.
package main

// The import statement is used to bring in other packages from the Go standard library.
// It allows you to access functions, types, and methods that are defined in other packages. 
import "fmt" // Importing fmt package for formatted I/O.

// Imports can also be done with an import block.
// You can search for packages at https://pkg.go.dev
import (
	"math" // Importing math package for mathematical functions.
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!") // Using fmt package to print to console.
	fmt.Println(math.Sqrt(16)) // Using math package to calculate square root.
	// "quote" is the package and "Go" is the function inside the package.
    	// Meaning you need to know the source code of the package and the function, or at least know how to use them.
    	// Package info: https://pkg.go.dev/rsc.io/quote/v4#Go
    	// Function source code: https://github.com/rsc/quote/blob/v4.0.1/quote.go#L22
	fmt.Println(quote.Go())
}

// To compile a stand alone executable, run: go build -o GoLozFun