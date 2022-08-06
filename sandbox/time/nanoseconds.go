package main

import (
	"fmt"
	"time"
)

func main(){
	// returns the number(int64) of nanoseconds since Jan 1 1970 UTC
	fmt.Println(time.Now().UnixNano())
}
