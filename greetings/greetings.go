package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting of the named person.
func Hello(name string) (string, error) {
	// if the name is empty throw an error
	if name == "" {
		return "", errors.New("Name can not be empty, provide a name.\n\n")
	}
	message := fmt.Sprintf("\nHello %v, welcome to the Go community!\n\n", name)

	return message, nil
}

// fmt.Sprintf

//   ~ a format string, similar to python format string literals:
//       + f string literal or more precisely literal string dot format:
// 		     - 'Hello %v, welcome to the Python community.format(message)'

// errors package

//   - a package that is part of the standard library
//   - used to handle and throw errors

// errors.New

//   - returns a new error with the passed in error message argument

// error handling with nil

//   - when error handling in Go you return nil for the error type if there is no error

// Functions Returning Multiple Values

//   - In Go functions can return multiple values instead of needing to use a tuple / Record
//   - the caller unpacks (Python) / destructures (TypeScript / Dart) the values
//     into variables on the recieving end

// multi return statements

//   - every return statement used must return the number of variables
//     specified in the return type head

// Return data-type head

//   - in Go you can ha ve two heads
//   - the first head specifies the set of parameters and their types
//   - the second head can be ommited if there is only 1 returned variable
//   - if there is more than 1 variable being returned a second head is required
//   - providing variable names within the return type head is optional (why add more optional code? K.I.S.M)
