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
	sages := map[string]string{"INDIA": "Gandhi", "USA": "MLK", "MEDIATTE": "Budhha", "LOVE": "Jesus"}
	err := tpl.Execute(os.Stdout, sages)
	printError(err)
}
