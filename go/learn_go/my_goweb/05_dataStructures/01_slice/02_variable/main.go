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
	myFamily := []string{"Dazzie", "Adam", "Hannah", "Speckles"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", myFamily)
	if err != nil {
		log.Fatalln(err)
	}
}
