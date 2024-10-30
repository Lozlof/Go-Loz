package lozgo

import (
	"errors"
	"log"
	"strconv"
)

// In Go, "error" is a predefined type.
// Takes a string input and two int's.
// The string is sent off to sanitizationIntVerify to be verified and converted to an int.
func SanitizationIntInRange(input string, minNum, maxNum int) (int, error) {
	// Set logging values.
	log.SetPrefix("lozgo/SanitizationIntInRange: ")
	log.SetFlags(0)

	// Attempts to convert "input" to an integer.
	// If it is converted, err will be nil.
	// convertedInt will be initialized to the integer returned by strconv.Atoi.
	convertedInt, err := strconv.Atoi(input)

	// Returns 0 and error if input cannot be converted to an integer.
	if err != nil {
		// Printf calls Output to print to the standard logger. 
		// Arguments are handled in the manner of fmt.Printf.
		// Try simplifying the error logging to remove any unintended values from the output.
		log.Printf("%q cannot be converted to an integer", input)
		return 0, errors.New("Error. See log output.")
	} else {
		// Returns the valid integer if it's in range.
		if convertedInt >= minNum && convertedInt <= maxNum {
			return convertedInt, nil
		} else {
			// The Errorf function lets us use formatting features to create descriptive error messages. 
			// In Go's fmt package, %q and %d are formatting verbs used to represent different types of data in formatted output.
			// If name = "bueller", then %q will output "bueller" (with quotes around the string).
			// In the example, id = 17, so %d will output 17 as a standard integer without any extra formatting.
			log.Printf("%d is not within the range %d - %d", convertedInt, minNum, maxNum)
			return 0, errors.New("Error. See log output")
		}
	}
}