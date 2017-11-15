package main
import (
	"text/template"
	"os"
	"fmt"
)
func main() {
	s1,_ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout,"header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)

	//t, _ :=template.ParseFiles("templates/layout.html")
	//fmt.Println(t.Name())
	//t.Execute(w, "Hello world")
}