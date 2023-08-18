package main

// TODO: ad comments from example

import (
	"fmt"
	"log"

	"example.com/greetings"
)

// Logger encapsulates logging errors.
type Logger struct {
}

// Error sets prefixes, flags, prints error and exits the program.
func (logger Logger) Error(err error, logger_prefix string) {

	// Sets the output prefix displayed to the console from the logger.
	log.SetPrefix(logger_prefix)

	// Sets the output flags for the logger.
	log.SetFlags(0)

	// Prints the error concatenating the logger prefix and exits the application with os.Exit(1)
	log.Fatal(err)
}

func main() {
	logger := Logger{}
	message, err := greetings.Hello("")
	if err != nil {
		logger.Error(err, "greetings: ")
	}
	fmt.Printf(message)
}

// Executing Go code as an Application

//   - declaring a package main lets the Go compiler know its an executable application
//   - the declared main package does not have to be called main.go
//   - there can only be one main package per module
