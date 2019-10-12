package main

import (
	"log"
	"os"
	"text/template"
)

func printError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	tpl, err := template.ParseFiles("./tpl.goHtml")
	printError(err)
	nf, err := os.Create("index.html")
	printError(err)
	err = tpl.Execute(nf, nil)
	printError(err)
}
