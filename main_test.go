package main

import (
	"os"
	"testing"
)

func TestGoTemplate(t *testing.T) {
	cases := []struct {
		name          string
		template      string
		resultantfile string
	}{
		{
			name:          "Test for ensuring template and resulting file are created",
			template:      "example.css.template",
			resultantfile: "example.css",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			// create the template
			createtemplate(templatepath)

			// create the file from the parsed template
			parse(templatepath, newfilepath)

			// Check both files were created
			_, err := os.Stat(c.template)
			os.IsNotExist(err)
			if err != nil {
				panic(err)
			}
			_, err = os.Stat(c.resultantfile)
			os.IsNotExist(err)
			if err != nil {
				panic(err)
			}

			// Delete each file after running test
			err = os.Remove(c.template)
			if err != nil {
				panic(err)
			}

			err = os.Remove(c.resultantfile)
			if err != nil {
				panic(err)
			}

		})
	}
}
