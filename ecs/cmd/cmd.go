package main

import (
	"fmt"
	"os"
	"text/template"

	_ "embed"
)

//go:embed system.tmpl
var systemTemplate string

func main() {
	d := struct {
		SystemName string
	}{
		SystemName: "TestSystem",
	}


	tmpl, err := template.New("myTemplate").Parse(systemTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	// Execute the template with the data
	err = tmpl.Execute(os.Stdout, d)
	if err != nil {
		panic(err)
	}
}
