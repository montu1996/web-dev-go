package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// Execute on initilization.
func init() {
	// Error checkign done inside must.
	tpl = template.Must(template.ParseGlob("./templates/*.goHtml"))
}

func printError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "01.goHtml", nil)
	printError(err)
	fmt.Println()
	err = tpl.ExecuteTemplate(os.Stdout, "02.goHtml", nil)
	printError(err)
	fmt.Println()
	err = tpl.ExecuteTemplate(os.Stdout, "03.goHtml", nil)
	printError(err)
	fmt.Println()
}
