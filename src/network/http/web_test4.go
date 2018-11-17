package main

import (
	"fmt"
	"net/http"
)

type Mymux struct {
}

func (mux *Mymux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		IndexFunc(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func IndexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world,%s !", r.URL.Path[:])
}


func main() {
	mux := &Mymux{}
	http.ListenAndServe(":8080", mux)
}
