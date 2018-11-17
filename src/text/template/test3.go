package main

import (
	"text/template"
	"os"
	"fmt"
)

func main(){

	s1,_ := template.ParseFiles("src/text/template/header.tmpl","src/text/template/content.tmpl","src/text/template/footer.tmpl")


	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()



}



