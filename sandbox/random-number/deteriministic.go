package main

import (
	"fmt"
	"math/rand"
)

func randNum() int {
	return rand.Intn(100)
}

func main(){
		fmt.Println("Random Number : ", randNum())
}
