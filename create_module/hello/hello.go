package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main(){
	// Set properts of the predefinded looger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// get a gretting message and print
	message, err := greetings.Hello("")
	// If an error was returned, print to the console and exit
	if err != nil {
		log.Fatal(err)
	}
	

	// If no error was returned, print the returned message to console
	fmt.Println(message)
}
