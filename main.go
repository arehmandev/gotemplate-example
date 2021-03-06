package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"k8s.io/helm/pkg/chartutil"
)

var (
	templatepath     = "example.css.template"
	templatecontents = "The text color is {{.textColor}} and the link color is {{.linkColorHover}}. Nestedkey: {{.testkey.testkeynested}}. Nested array: {{index .testkey.testkeylist 0}}"
	newfilepath      = "example.css"
	variablesfile    = "config.yml"
)

func main() {

	// create the template
	createtemplate(templatepath, templatecontents)

	// create the file from the parsed template
	parse(templatepath, newfilepath, variablesfile)

}

func createtemplate(createtemplatepath, createtemplatecontents string) {
	// create the template
	f, _ := os.Create(createtemplatepath)
	f.Write([]byte(createtemplatecontents))
	f.Close()
}

func parse(parsedtemplate, resultingfile, configfile string) {
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

	// Taking yaml values from file, thanks helm packages!
	readvalues, err := chartutil.ReadValuesFile(configfile)
	if err != nil {
		log.Print("executing template error: ", err)
		return
	}

	config := readvalues.AsMap()

	// fmt.Println(config.AsMap()) // uncomment to see the values printed as a map

	err = t.Execute(f, config)
	if err != nil {
		log.Print("executing template error: ", err)
		return
	}

	fmt.Println("Template and Resultingfile produced:")
	fmt.Printf("%v and %v were created \n", parsedtemplate, resultingfile)

	// fmt.Println(config["testkey"])

	// Uncomment to have the stdout (the contents of the newfile) on your console
	// fnew, _ := os.Open(resultingfile)
	// io.Copy(os.Stdout, fnew)
	// f.Close()
}
