package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	MOdel        string
	Doors        string
}

type items struct {
	Wisdeom   []sage
	Transport []car
}

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
		Motto: "Don't overthink, just let it go.",
	}

	c1 := car{
		Manufacturer: "TOYOTA",
		MOdel:        "Toyota Camry",
		Doors:        "4",
	}

	gandhi := sage{
		Name:  "GANDHI",
		Motto: "Happiness is when what you think, what you say, and what you do are in harmony.",
	}

	c2 := car{
		Manufacturer: "BMW",
		MOdel:        "BMW X5",
		Doors:        "4",
	}

	sages := []sage{
		buddha,
		gandhi,
	}

	cars := []car{
		c1,
		c2,
	}

	item := items{
		Wisdeom:   sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, item)

	printError(err)

	fmt.Println()
}
