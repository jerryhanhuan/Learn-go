package main

import (
	"fmt"
	"net/http"
	"strings"
	"html/template"
	"log"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		//t, _ := template.ParseFiles("src/network/http/login.gtpl")
		//t.Execute(w, nil)
		template.Must(template.ParseFiles("src/network/http/login.gtpl")).Execute(w, r)

	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断

		// method 1
		/*
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		*/
		//method 2 FormValue 会自动调用 ParseForm(),FormValue 只会返回同名参数中的第一个，若参数不存在，则返回空字符串
		fmt.Println("username:",r.FormValue("username"))
		fmt.Println("password:",r.FormValue("password"))

		fmt.Println("username2:", r.Form["username"])
		fmt.Println("password2:", r.Form["password"])

		fmt.Println("username3",r.Form.Get("username"))
		fmt.Println("password3",r.Form.Get("password"))
		fmt.Println("login_method:",r.Form.Get("login_method"))
		fmt.Println("gender:",r.Form.Get("gender"))
		fmt.Println("interest:",r.Form.Get("interest"))
	}
}
func main() {
	t := time.Date(2018,time.November,7,15,25,45,0,time.UTC)
	fmt.Printf("Go launched at %s.\n",t.Local())
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
