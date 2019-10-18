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
	g1 := struct {
		Score1 int
		Score2 int
	}{
		7, 9,
	}

	err := tpl.Execute(os.Stdout, g1)

	printError(err)

	fmt.Println()

}
