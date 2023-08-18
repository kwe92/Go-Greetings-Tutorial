package main

// TODO: ad comments from example

import (
	"fmt"

	"example.com/greetings"
	"example.com/logger"
)

func main() {
	logger := logger.Logger{}
	message, err := greetings.Hello("")
	if err != nil {
		logger.Error(err, "\ngreetings: ")
	}
	fmt.Printf(message)
}

// Executing Go code as an Application

//   - declaring a package main lets the Go compiler know its an executable application
//   - the declared main package does not have to be called main.go
//   - there can only be one main package per module
