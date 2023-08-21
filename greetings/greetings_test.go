package greetings

import (
	"regexp"
	"testing"

	// "fmt"
	utils "example.com/utility"
)

// Given: Hello function is passed a non-empty string for the name property
// When: Hello function executes
// Then: the name should exist in the returned message and there should be no error message
func TestHelloName(test *testing.T) {

	name := "Kweayon"

	// a regular expression to parse the name from the message
	want := regexp.MustCompile(`\b` + name + `\b`)

	// calls Hello and unpacks `destructures` its results
	msg, err := Hello(name)

	// if the message does not contain the name or if there was an error fail the test.
	if !want.MatchString(msg) || utils.HasError(err) {
		test.Fatalf(`Hello("Kweayon") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}

// Given: Hello function is passed a empty string
// When: Hello function is executed
// Then: an empty string and a error message should be returned

func TestHelloEmpty(test *testing.T) {
	msg, err := Hello("")

	// if the message is not empty or if there was no error thrown
	if msg != "" || err == nil {
		test.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// ------------------------SUMMARY------------------------ //

// TODO: draw test routine function signature on paper, break it down

// What is Automated Testing?

//   - Implementation of an automation tool to execute test cases

// Why Write Automated Tests?

//       + decrease running into bugs as your code base expands
//       + free up time from testers to worek on more pressing matters
//       + write bug free code
//       + increase application quality, scalability ansd test coverage

// Test Coverage

//   - The total number of functions over the number of functions with tests written

// testing package

//   - part of the Go standard library
//   - provides automated testing of Go packages
//   - similar to the `test` package that is part of the
//     flutter standard library
//   - allows you developers write unit tests for go packages in the form
//     of test routines and test suites

// go test CLI command

//   - used in conjunction with the testing package
//   - the functions in files suffixed with _test.go
//     are considered as test routines and test suites
//   - and go test executes them as such
//   - analogous to the flutter test command
//   - go test -v (added verboseness to the output test message)

// Test Routine

//   - a single unit test
//   - the signature of a unit test

// Test Suite

//   - a set of one or more test routines

// Exception Handling in GO

//   - there are no exceptions in Go
//   - implying that exception handling
//     keywords try except finally do not exist
//   - which means fatal errors which are errors
//     that usually would be caught by the catch exception handler
//     must be handled manually within your code

// Signaling Failure

//   - the test routine callback takes a testing structure
//     as an argument, which allows you to handle fatal errors

// Struct Parameters

//   - structures that are passed as arguments
//   - to functions by reference `pointer`

// TODO: figure out the below issue
// Test Suite Package Name
//?   - test suite package name needs to match the package you want to test or is it the file name?
