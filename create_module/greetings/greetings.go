package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that returns the name in varible message
	message := fmt.Sprintf("Hi, %v. Welcome!", name) // := decleare & initialize using the value on the right
// %v => value 
	return message, nil
}
