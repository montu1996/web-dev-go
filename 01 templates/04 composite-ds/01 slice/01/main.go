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
		log.Fatalln(err)
	}
}

func main() {
	sages := []string{"Gandhi", "MLK", "Budhha", "Jesus"}
	err := tpl.Execute(os.Stdout, sages)
	printError(err)
}
