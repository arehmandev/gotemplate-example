package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

var (
	templatepath     = "example.css.template"
	templatecontents = "The text color is {{.textColor}} and the link color is {{.linkColorHover}} {{.testkey.testkeynested}}"
	newfilepath      = "example.css"
)

type values map[string]interface{}

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
	config, err := readValuesFile("config.yml")
	if err != nil {
		log.Print("executing template error: ", err)
		return
	}

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

func readValuesFile(filename string) (values, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return readValues(data)
}

func readValues(data []byte) (vals values, err error) {
	err = yaml.Unmarshal(data, &vals)
	if len(vals) == 0 {
		vals = values{}
	}
	return
}
