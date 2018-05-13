package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42) // 传入数据执行
	if err != nil {
		log.Fatalln(err)
	}
}