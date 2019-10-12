package main

import (
	"fmt"
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
	names := []string{
		"Montu",
		"Mitesh",
		"PK",
	}

	err := tpl.Execute(os.Stdout, names)
	printError(err)
	fmt.Println()
}
