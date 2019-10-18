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

// Person class.
type Person struct {
	Name string
	Age  string
}

// doubleZero class.
type doubleZero struct {
	Person
	IsLicenceToKill bool
}

func main() {

	p1 := Person{
		Name: "Montu",
		Age:  "23",
	}

	dz1 := doubleZero{
		Person:          p1,
		IsLicenceToKill: true,
	}

	err := tpl.Execute(os.Stdout, dz1)
	printError(err)
	fmt.Println()
}
