package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
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

	buddha := sage{
		Name:  "BUDDHA",
		Motto: "Don't overthink, just let it go",
	}
	gandhi := sage{
		Name:  "GANDHI",
		Motto: "Happiness is when what you think, what you say, and what you do are in harmony.",
	}
	jesus := sage{
		Name:  "JESUS",
		Motto: "Do not let your hearts be troubled. Trust in God; trust also in me.",
	}

	sages := []sage{
		buddha,
		gandhi,
		jesus,
	}

	err := tpl.Execute(os.Stdout, sages)
	printError(err)
	fmt.Println()

}
