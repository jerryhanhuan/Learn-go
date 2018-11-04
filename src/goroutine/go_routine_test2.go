package main

import (
	"time"
	"fmt"
)

func main() {

	go func() {
		time.Sleep(time.Millisecond*100)
		fmt.Println("1")
	}()

	go func() {
		time.Sleep(time.Millisecond*200)
		fmt.Println("2")
	}()

	go func() {
		time.Sleep(time.Millisecond*300)
		fmt.Println("3")
	}()


	time.Sleep(time.Millisecond*400)

}


