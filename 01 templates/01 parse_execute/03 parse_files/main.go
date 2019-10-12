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
	tpl, err := template.ParseFiles("01")
	printError(err)
	err = tpl.Execute(os.Stdout, nil)
	fmt.Println()
	printError(err)
	tpl, err = template.ParseFiles("02", "03")
	printError(err)
	err = tpl.ExecuteTemplate(os.Stdout, "03", nil)
	printError(err)
	fmt.Println()
	err = tpl.ExecuteTemplate(os.Stdout, "02", nil)
	printError(err)
	fmt.Println()
	err = tpl.Execute(os.Stdout, nil)
	printError(err)
}
