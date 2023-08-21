package utility

import (
	"math/rand"
)

// RandInt returns an integer from [0, exclusive_range)
func RandInt(exclusive_range int) (random_integer int) {
	return rand.Intn(exclusive_range)
}

// IsEmpty is a predicate checking if a string is empty.
func IsEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false

}

// HasError is a predicate checking if an error was returned.
func HasError(err error) bool {
	if err != nil {
		return true
	}
	return false
}
