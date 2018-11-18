package main

import (
	"os"
	"fmt"
	"net"
)

/*
type TCPAddr struct {
    IP IP
    Port int
}

go run test2.go  www.baidu.com:80

go run test2.go  192.1.2.234:1818

 */

func main(){
	//resolve ip
	if len(os.Args)!=2 {
		fmt.Printf("Usage :%s host\n",os.Args[0])
		os.Exit(1)
	}
	host := os.Args[1]

	ip,err := net.ResolveTCPAddr("tcp",host)
	checkError(err)
	fmt.Println(ip)
	fmt.Printf("ip:%s port:%d \n",ip.IP,ip.Port)
}


func checkError(err error) {
	if err!=nil {
		fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
		os.Exit(1)
	}
}