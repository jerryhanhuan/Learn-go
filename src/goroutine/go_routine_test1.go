package main

import (
	"fmt"
	"runtime"
)

/*

runtime.Gosched函数的作用是让当前正在运行的Goroutine（这里是运行main函数的那个Goroutine）暂时“休息”一下，
而让Go运行时系统转去运行其它的Goroutine（这里是与go fmt.Println("Go!")对应并会封装fmt.Println("Go!")的那个Goroutine）。
如此一来，我们就更加精细地控制了对几个Goroutine的运行的调度。
 */

func main(){

	go fmt.Printf("Hello World")
	runtime.Gosched()

}
