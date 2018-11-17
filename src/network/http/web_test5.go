package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method::", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%X", h.Sum(nil))
		//渲染 HTML
		template.Must(template.ParseFiles("src/network/http/upload.html")).Execute(w, token)
	} else {
		fmt.Println("Start upload")
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

}

func main() {

	//注册路由
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("http.ListenAndServe", err)
	}
}
