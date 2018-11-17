package main

import (
	"html/template"
	"os"
)

type person struct{
	Name string
	Email string
}

func main(){

	t := template.New("FileName example")
	t.Parse("Hello {{.Name}} {{.Email}}")
	p := person{Name:"Yudf",Email:"yudf@keyou.cn"}
	t.Execute(os.Stdout,p)  //os.Stdout 实现了 io.Writer 接口
}



