package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

// Course class.
type Course struct {
	Number, Name, Units string
}

// Semester class.
type Semester struct {
	Term    string
	Courses []Course
}

// Year class.
type Year struct {
	Fall, Spring, Summer Semester
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
	y := Year{
		Fall: Semester{
			Term: "Fall",
			Courses: []Course{
				Course{"A", "Description of A", "100"},
				Course{"B", "Description of B", "100"},
			},
		},
		Spring: Semester{
			Term: "Spring",
			Courses: []Course{
				Course{"C", "Description of C", "100"},
				Course{"D", "Description of D", "100"},
			},
		},
	}
	err := tpl.Execute(os.Stdout, y)
	printError(err)
	fmt.Println()
}
