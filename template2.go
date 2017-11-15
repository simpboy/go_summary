package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email string  //不报错---,输出空
}

func main() {
	t := template.New("fieldname example")
	t,_ = t.Parse("Hello {{.UserName}}---{{.email}}")
	p := Person{UserName:"Astaxie",email:"astaxie@163.com"}
	t.Execute(os.Stdout,p)
}
