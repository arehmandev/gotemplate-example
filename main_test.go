package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestGoTemplate(t *testing.T) {
	cases := []struct {
		name                  string
		template              string
		templatefilecontents  string
		resultantfile         string
		resultantfilecontents string
	}{
		{
			name:                  "Test for ensuring template and resulting file are created",
			template:              templatepath,
			templatefilecontents:  templatecontents,
			resultantfile:         newfilepath,
			resultantfilecontents: "The text color is #abcdef and the link color is #ffaacc",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			// create the template
			createtemplate(templatepath, templatecontents)

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

			// Check the contents of both files is correct

			//templatefilecontents
			b1, err := ioutil.ReadFile(c.template)
			if err != nil {
				fmt.Print(err)
			}

			result1 := strings.Replace(string(c.templatefilecontents), "\u00a0", " ", -1)
			if string(b1) != result1 {
				t.Errorf("got: %#v\nwant: %#v\n", string(b1), result1)
			}

			//resultantfilecontents
			b2, err := ioutil.ReadFile(c.resultantfile)
			if err != nil {
				fmt.Print(err)
			}

			result2 := strings.Replace(string(c.resultantfilecontents), "\u00a0", " ", -1)
			if string(b2) != result2 {
				t.Errorf("got: %#v\nwant: %#v\n", string(b2), result2)
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
