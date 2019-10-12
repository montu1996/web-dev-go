package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./tpl.goHtml"))
}

func printError(err error) {
	if err != nil {
		log.Fatalln("./tpl.goHtml")
	}
}

func main() {
	err := tpl.Execute(os.Stdout, "Montu")
	printError(err)
}
