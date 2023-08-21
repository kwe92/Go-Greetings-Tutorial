package greetings

import (
	"errors" // Standard Library: Implements functions to manipulate errors.
	"fmt"    // Standard Library: Implements formatted input/output functions with simple verbs like `%v`

	utils "example.com/utility"
)

// Hello returns a greeting of the named person.
func Hello(name string) (string, error) {
	// if the name is empty throw an error

	if utils.IsEmpty(name) {
		return "", errors.New("Name can not be empty, provide a name.\n\n")
	}
	message := fmt.Sprintf(randomGreeting(), name)
	// message := fmt.Sprintf(randomGreeting())

	return message, nil
}

// Hellos returns a greeting for every person in the array of people.
func Hellos(people []string) (map[string]string, error) {
	greetings := make(map[string]string)
	for _, name := range people {
		greeting, err := Hello(name)
		if utils.HasError(err) {
			return nil, err
		}
		greetings[name] = greeting
	}
	return greetings, nil
}

// randomGreeting returns a random greeting.
func randomGreeting() string {
	// slice consisting of greetings as elements.
	greetings := []string{
		"\nHi, %v. Welcome!\n\n",
		"\nGreat to see you, %v!\n\n",
		"\nHail, %v! Well met!\n\n",
	}

	sliceLength := len(greetings)

	return greetings[utils.RandInt(sliceLength)]

}

// ?--------------------Summary--------------------?//

// fmt.Sprintf

//   ~ a format string, similar to pythons format string literals:
//       + f string literal or literal string dot format:
// 		     - 'Hello %v, welcome to the Python community'.format(message)

// errors package

//   - a package that is part of the standard library
//   - used to handle and throw errors

// errors.New

//   - returns a new error as a string with the passed in error message argument

// error handling with nil

//   - when error handling in Go you return nil for the error type if there is no error

// Functions Returning Multiple Values

//   - In Go functions can return multiple values instead of needing to use a tuple or Record
//   - the caller unpacks (Python) / destructures (TypeScript / Dart) the values
//     into variables on the recieving end

// multi return statements

//   - every return statement used must return the number of variables
//     specified in the return type head

// Return data-type head

//   - in Go you can have two heads
//   - the first head specifies the set of parameters and their types
//   - the second head can be ommited if there is only 1 returned variable
//   - if there is more than 1 variable being returned a second head is required
//   - providing variable names within the return type head is optional (why add more optional code? K.I.S.M)

// Array (fixed-length Iterable) | Slice (Growable Iterable)

// Slices are abstractions of arrays

// for range GO

//   - similar to pythons for in enumerate
//   - using range automatically includes the enumerated index
//   - if you do not need the index you can use the blank identifier _ like in most other languages
//   - e.g.
//       - Go: for index, name := range names {...}
//       - Python: for index, name in enumerate(names):
//                     ...

// make function

//   - similar to the new keyword in other languages
//   - instead of returning a reference `pointer` to the created object it returns the object
