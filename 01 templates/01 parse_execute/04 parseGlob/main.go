package main

import (
	"fmt"
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
	tpl, err := template.ParseGlob("./templates/*.goHtml")
	printError(err)
	err = tpl.ExecuteTemplate(os.Stdout, "01.goHtml", nil)
	printError(err)
	fmt.Println()
	err = tpl.ExecuteTemplate(os.Stdout, "02.goHtml", nil)
	printError(err)
	fmt.Println()
	err = tpl.ExecuteTemplate(os.Stdout, "03.goHtml", nil)
	printError(err)
	fmt.Println()
}
