package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"text/template"
)

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqrt(x float64) float64 {
	return math.Sqrt(float64(x))
}

var fm = template.FuncMap{
	"dbl":  double,
	"sqr":  square,
	"sqrt": sqrt,
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
	err := tpl.Execute(os.Stdout, 100)
	printError(err)
	fmt.Println()
}
