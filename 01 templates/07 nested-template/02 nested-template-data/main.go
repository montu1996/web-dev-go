package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.goHtml"))
}

func printError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.goHtml", "MONTU")
	printError(err)
	fmt.Println()
}
