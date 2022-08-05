package main

import (
	"fmt"

	"example.com/greetings"
)

func main(){
	// get a gretting message and print
	message := greetings.Hello("Burt")

	fmt.Println(message)
}
