package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	// Get Current Directory.
	currDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}

	// Read File from current Directory. ( only works on build case )
	path, err := filepath.Abs(currDir + "/tpl.goHtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Parse File
	tpl, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute template into standard output
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
