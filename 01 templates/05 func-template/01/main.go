package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("tpl.goHtml").Funcs(fm).ParseFiles("./tpl.goHtml"))
}

func printError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	buddha := sage{
		Name:  "BUDDHA",
		Motto: "Don't ovethink, just let it go",
	}
	montu := sage{
		Name:  "MONTU",
		Motto: "Keep Learning",
	}

	sages := []sage{
		buddha,
		montu,
	}

	err := tpl.Execute(os.Stdout, sages)

	printError(err)

	fmt.Println()

}
