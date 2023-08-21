package logger

import (
	"log"
)

// Error sets prefixes, flags, prints error and exits the program.
func Error(err error, logger_prefix string) {

	// Sets the output prefix displayed to the console from the logger.
	log.SetPrefix(logger_prefix)

	// Sets the output flags `time stamps` for the logger.
	log.SetFlags(0)

	// Prints the error concatenating the logger prefix and exits the application with os.Exit(1).
	log.Fatal(err)
}

// Implementing Methods For a Struct

//   - you can not implement methods directly in a struct
//   ~ to implement a method for a struct you must:

//       + define a callback outside of the struct that takes
//         the struct as an argument
//       + the callback acts as a method to any instances of the passed in struct
