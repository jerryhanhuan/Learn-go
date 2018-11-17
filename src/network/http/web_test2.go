package main

import (
	"net/http"
	"fmt"
)

func sayHello(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer,"Hello to the mutex,%s!",request.URL.Path[1:])
}



func main(){

	//创建一个多路复用器
	mux := http.NewServeMux()
	mux.HandleFunc("/",sayHello)
	server := &http.Server{
	Addr:"127.0.0.1:8080",
	Handler:mux,
	}
	server.ListenAndServe()

}

