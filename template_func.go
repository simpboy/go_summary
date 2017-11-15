package main

import (
	"fmt"
	"strings"
	"html/template"
	"os"
)

type Friend_f struct{
	Fname string
}

type Person_f struct {
	UserName string
	Emails []string
	Friends []*Friend_f
}


func main() {
	f1 := Friend_f{"minux_xu"}
	f2 := Friend_f{"xushiwei"}

	t := template.New("template test")
	t.Funcs(template.FuncMap{"emailDeal":EmailDealWith})
	t,_ = t.Parse(`
		hello {{.UserName}}!
		{{range .Emails}}
		an emails {{.|emailDeal}}
		{{end}}
		{{with .Friends}}
		{{range .}}
			my friend name is {{.Fname}}
		{{end}}
		{{end}}
	`)

	p := Person_f{
			UserName:"astaxie",
			Emails:[]string{"astaxie@163.com","astaxie@gmail.com"},
			Friends:[]*Friend_f{&f1,&f2}}
	t.Execute(os.Stdout,p)



}
func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return (substrs[0] + " at " + substrs[1])
}