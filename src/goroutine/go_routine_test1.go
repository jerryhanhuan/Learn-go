package main

import (
	"fmt"
	"runtime"
)

func main(){

	go fmt.Printf("Hello World")
	runtime.Gosched()

}
