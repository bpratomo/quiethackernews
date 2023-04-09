package templateparse

import (
	"fmt"
	"html/template"
)

func Render() *template.Template {
	// Load the HTML file containing the template
	templateFile := "templateparse/base.html"
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return nil
	}
    return tmpl

}
