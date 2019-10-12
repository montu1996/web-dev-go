package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func dateMonthYear(t time.Time) string {
	return t.Format("02-Jan-2006")
}

var fm = template.FuncMap{
	"fdateDMY": dateMonthYear,
}

func init() {
	tpl = template.Must(template.New("tpl.goHtml").Funcs(fm).ParseFiles("./tpl.goHtml"))
}

func printError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	err := tpl.Execute(os.Stdout, time.Now())
	printError(err)
	fmt.Println()
}
