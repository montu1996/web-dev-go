package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type user struct {
	Name  string
	Motto string
	Admin bool
}

func printError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./tpl.goHtml"))
}

func main() {
	buddha := user{
		Name:  "Buddha",
		Motto: "Don't overthink, just let it go.",
		Admin: true,
	}

	gandhi := user{
		Name:  "Gandhi",
		Motto: "Bo the Change.",
		Admin: false,
	}

	users := []user{buddha, gandhi}

	err := tpl.Execute(os.Stdout, users)

	printError(err)

	fmt.Println()

}
