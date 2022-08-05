package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that returns the name in varible message
	message := fmt.Sprintf("Hi, %v. Welcome!", name) // := decleare & initialize using the value on the right
// %v => value 
	return message
}
