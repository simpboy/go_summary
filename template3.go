package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person_3 struct {
	UserName string
	Emails	[]string
	Friends	[]*Friend
}
func main() {
	f1 := Friend{Fname:"minxu.ma"}
	f2 := Friend{Fname:"xushiwei"}

	t := template.New("fieldname example")
	t,_ = t.Parse(`Hello {{.UserName}}!
	{{range .Emails}}
	an email:{{.}}{{end}}
		{{with .Friends}}
		{{range .}}
			my friend name is {{.Fname}}
		{{end}}
		{{end}}
	`)

	P := Person_3{UserName:"Astaxie",
		Emails : []string{"astaxie@beego.me","astaxie@gmail.com"},
		Friends : []*Friend{&f1,&f2}}
	t.Execute(os.Stdout,P)
}
