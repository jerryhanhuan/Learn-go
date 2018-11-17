package main

import (
	"os"
	"fmt"
	"net"
)

func main(){
		if len(os.Args) !=2{
			fmt.Printf("Usage:%s ip-addr \n",os.Args[0])
			os.Exit(1)
		}
		ip := net.ParseIP(os.Args[1])
		if ip !=nil {
			fmt.Printf("ip is %s",ip.String())
		}else {
			fmt.Printf("Invalid address")
		}
		os.Exit(0)
}

