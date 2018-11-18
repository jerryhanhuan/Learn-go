package main

import (
	"os"
	"fmt"
	"strconv"
	"net"
	"time"
)

func main(){

	if len(os.Args)!=2 {
		fmt.Printf("Usage %s port \n",os.Args[0])
		os.Exit(1)
	}
	var ip string
	port,_ := strconv.Atoi(os.Args[1])
	fmt.Printf("port:%d \n",port)
	ip = fmt.Sprintf(":%d",port)
	fmt.Printf("ip is %s\n",ip)
	svraddr,err := net.ResolveTCPAddr("tcp",ip)
	checkError2(err)

	//listen
	listener,err := net.ListenTCP("tcp",svraddr)
	checkError2(err)
	fmt.Println("listen ok")
	for{
		//acept
		fmt.Println("waiting for client connection")
		conn,err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr,"listener Accept err %s ",err.Error())
			continue
		}
		fmt.Printf("accept ok  %s form %s",conn.LocalAddr().String(),conn.RemoteAddr().String())
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}

}


func checkError2(err error) {
	if err!=nil {
		fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
		os.Exit(1)
	}
}