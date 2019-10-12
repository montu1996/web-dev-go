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
	data := struct {
		Words []string
		Lname string
	}{
		[]string{"1", "2", "3"},
		"Thakor",
	}

	err := tpl.Execute(os.Stdout, data)
	printError(err)
	fmt.Println()
}
