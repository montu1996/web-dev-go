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

type Person struct {
	Name string
	Age  string
}

func main() {

	p1 := Person{
		Name: "Montu",
		Age:  "23",
	}

	err := tpl.Execute(os.Stdout, p1)
	printError(err)
	fmt.Println()
}
