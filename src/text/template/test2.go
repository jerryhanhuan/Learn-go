package main

import (
	"html/template"
	"os"
)

type friend struct{
	Fname string
}
type people struct{
	Name string
	Emails []string
	Friends []*friend
}

func main(){

	f1 := friend{Fname:"mohz"}
	f2 := friend{Fname:"yexb"}
	f3 := friend{Fname:"zhaoyp"}
	email := []string{"yudf@keyou.cn","dengfeng2012@163.com","dengfeng@google.com"}
	p1 := people{"yudf",email,[]*friend{&f1,&f2,&f3}}

	t := template.New("template test")
	t.Parse("hello {{.Name}} ! your email {{range .Emails}}{{.}}{{end}} your friends {{with .Friends}}{{range .}}{{.Fname}} {{end}}{{end}}")
	t.Execute(os.Stdout,p1)


}


