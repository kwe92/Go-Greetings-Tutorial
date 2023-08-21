package main

// TODO: ad comments from example

import (
	"fmt"

	"example.com/greetings"
	log "example.com/logger"
	utils "example.com/utility"
)

func main() {
	people := []string{
		"",
		"Kweayon",
		"Ronin",
		"Devin",
		"Lillia",
		"James",
	}

	randomPerson := people[utils.RandInt(len(people))]

	logError := log.Logger{}.Error
	message, err := greetings.Hello(randomPerson)
	messages, err := greetings.Hellos(people[1:5])
	if err != nil {
		logError(err, "\ngreetings: ")
	}
	fmt.Printf(message)
	fmt.Println(messages)
}

// Executing Go code as an Application

//   - declaring a package main lets the Go compiler know its an executable application
//   - the declared main package does not have to be called main.go
//   - there can only be one main package per module
