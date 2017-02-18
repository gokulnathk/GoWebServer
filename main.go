package main

import (
	"GoWebServer/controllers"
	"net/http"
	"os"
	"text/template"
)

func main() {
	templates := populateTemplates()

	controllers.Register(templates)

	http.ListenAndServe(":8800", nil)
}

func populateTemplates() *template.Template {
	// Creating new template with name templates
	result := template.New("templates")
	// setting path for the templates location (folder name)
	templatePath := "templates"
	// reading the template folder content from the path
	templateFolder, _ := os.Open(templatePath)
	// close the template folder once read complete using the defer key
	defer templateFolder.Close()

	// read template folder's content will be a slice
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)

	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, templatePath+"/"+pathInfo.Name())
		}
	}

	/**
	  parse the files present in the template paths, ParseFiles will accept a variadic
	  parameter adding ... to the end of the slice variable will convert the slice to comma separated strings
	*/
	result.ParseFiles(*templatePaths...)
	return result
}
