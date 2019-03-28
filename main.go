package main

import (
	"bytes"
	"encoding/json"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kylegrantlucas/resume-exporter/models"
	"github.com/kylegrantlucas/resume-exporter/templates"
	flag "github.com/ogier/pflag"
)

func main() {
	log.SetFlags(0)
	var infile, outfile, templateName string

	flag.StringVarP(&infile, "infile", "i", "resume.json", "the file to read in")
	flag.StringVarP(&outfile, "outfile", "o", "", "the file to output to")
	flag.StringVarP(&templateName, "template", "t", "classic", "the template to use")
	flag.Parse()

	// Load the resume json
	data, err := ioutil.ReadFile(infile)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatal(err)
	}

	// Unmarshal the JSON
	resume := models.Resume{}
	err = json.Unmarshal(data, &resume)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatal(err)
	}

	// Setup the Tex template
	tmpl, err := template.New(infile).Delims("[[", "]]").Parse(templateMap[templateName])
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatal(err)
	}

	// Exectute the template
	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, resume)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatal(err)
	}

	// Replace some strings that get escaped by acctident
	texString := buf.String()
	texString = strings.Replace(texString, "&amp;", "+", -1)
	texString = html.UnescapeString(texString)

	// Print out to stdout or file based on flags
	if outfile != "" {
		out, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.SetFlags(log.Lshortfile)
			log.Fatal(err)
		}

		_, err = out.Write([]byte(texString))
		if err != nil {
			log.SetFlags(log.Lshortfile)
			log.Fatal(err)
		}
	} else {
		log.Print(texString)
	}

	if classMap[templateName] != "" {
		out, err := os.OpenFile(templateName+".cls", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.SetFlags(log.Lshortfile)
			log.Fatal(err)
		}
		defer out.Close()

		out.Write([]byte(classMap[templateName]))
	}
}

var templateMap = map[string]string{
	"classic": templates.Classic,
	"modern":  templates.Modern,
}

var classMap = map[string]string{
	"modern": templates.ModernClass,
}
