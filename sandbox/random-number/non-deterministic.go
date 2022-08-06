package main 

import (
	"fmt"
	"math/rand"
)

func randomNumber() int64 {
	return rand.Seed(42)
}

func main (){
	fmt.Println("Random Number : ", randomNumber())
}
