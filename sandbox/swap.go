package main

import "fmt"

func swap(x, y string) (string, string){

	return y,x
}

func main(){

	a, b := swap("hello", "world")

	fmt.Println("No Swap : hello, world")

	fmt.Println("first swap : ", a,b)

	a, b = swap(a,b)
	fmt.Println("second swap  : ", a,b)
}
