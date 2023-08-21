package greetings

import (
	"regexp"
	"testing"

	// "fmt"
	utils "example.com/utility"
)

func TestHelloName(test *testing.T) {
	name := "Kweayon"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(name) || utils.HasError(err) {
		test.Fatalf(`Hello("Kweayon") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}

func TestHelloEmpty(test *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		test.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// ------------------------SUMMARY------------------------ //

// Test Suite Package Name
//?   - test suite package name needs to match the package you want to test or is it the file name?

// Test Routine

//   - a single unit test

// Test Suite

//   - s set of test routines

// Struct Parameters

//   - structures that are passed as arguments
//   - to functions by reference `pointer`
