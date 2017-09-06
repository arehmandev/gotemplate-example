package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var (
	templatepath     = "example.css.template"
	templatecontents = "The text color is {{.textColor}} and the link color is {{.linkColorHover}}"
	newfilepath      = "example.css"
)

func main() {
	// create the template
	createtemplate(templatepath, templatecontents)

	// create the file from the parsed template
	parse(templatepath, newfilepath)

}

func createtemplate(createtemplatepath, createtemplatecontents string) {
	// create the template
	f, _ := os.Create(createtemplatepath)
	f.Write([]byte(createtemplatecontents))
	f.Close()
}

func parse(parsedtemplate, resultingfile string) {
	t, err := template.ParseFiles(parsedtemplate)
	if err != nil {
		log.Println("parsing file error", err)
		return
	}

	f, err := os.Create(resultingfile)
	if err != nil {
		log.Println("creating file error: ", err)
		return
	}

	// A sample config
	config := map[string]string{
		"textColor":      "#abcdef",
		"linkColorHover": "#ffaacc",
	}

	err = t.Execute(f, config)
	if err != nil {
		log.Print("executing template error: ", err)
		return
	}

	fmt.Println("Template and Resultingfile produced:")
	fmt.Printf("%v and %v were created \n", parsedtemplate, resultingfile)

	// Uncomment to have the stdout (the contents of the newfile) on your console
	// fnew, _ := os.Open(resultingfile)
	// io.Copy(os.Stdout, fnew)
	// f.Close()
}
