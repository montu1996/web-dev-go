package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

// Page class.
type Page struct {
	Title   string
	Heading string
	Input   string
}

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
	home := Page{
		Title:   "Nothing Escaped",
		Heading: "Nothing Escaped with text/template",
		Input:   "<script>alert('YOOOO!!!')</script>",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.goHtml", home)
	printError(err)
	fmt.Println()
}
