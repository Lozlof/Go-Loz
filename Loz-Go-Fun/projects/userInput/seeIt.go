package main

// Use: go get -u github.com/Lozlof/lozgo to update to the newest lozgo repo.
import (
	"github.com/Lozlof/lozgo"
	"fmt"
)

func main() {
	// var userInput string

	var numStr string = "15"
	var numMin int = 5
	var numMax int = 20

	Number, err := lozgo.SanitizationIntInRange(numStr, numMin, numMax)

	fmt.Println("Number: ",Number)
	fmt.Println("err: ",err)
}