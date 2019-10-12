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

// Sage struct.
type Sage struct {
	Name  string
	Motto string
}

func main() {
	err := tpl.Execute(os.Stdout, Sage{Name: "Budhha", Motto: "Don't overthink, Just let it go"})
	printError(err)
	fmt.Println()
}
