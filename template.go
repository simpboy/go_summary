package main

import (
	"os"
	"log"
	"html/template"
)

func main() {
	//const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	//t := template.Must(template.New("escape").Parse(templ))
	//var data struct {
	//	A string        // untrusted plain text
	//	B template.HTML // trusted HTML
	//}
	//data.A = "<b>Hello!</b>"
	//data.B = "<b>Hello!</b>"
	//if err := t.Execute(os.Stdout, data); err != nil {
	//	log.Fatal(err)
	//}

	/**
	生成模板的输出需要两个处理步骤。第一步是要分析模板并转为内部表示，然后基于指定的输入执行模板。
	1、Parse
	2、Excute
	 */
	const templ = `<p>A:{{.A}}</p> <p>B:{{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data  struct{
		A string
		B template.HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout,data);err!=nil{
		log.Fatal(err)
	}




}
