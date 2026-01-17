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
	"time"

	"github.com/kylegrantlucas/resume-exporter/models"
	"github.com/kylegrantlucas/resume-exporter/templates"
	"github.com/kylegrantlucas/resume-exporter/validate"
	flag "github.com/ogier/pflag"
)

// escapeTypst escapes special Typst characters in content
// @ is interpreted as citation reference, so we escape it
func escapeTypst(s string) string {
	s = strings.ReplaceAll(s, "@", "\\@")
	return s
}

// formatDate converts ISO date strings to human-readable format
// "2022-04-01" -> "Apr 2022", "2022" -> "2022"
func formatDate(s string) string {
	if s == "" {
		return ""
	}

	// Try full date format: 2022-04-01
	if t, err := time.Parse("2006-01-02", s); err == nil {
		return t.Format("Jan 2006")
	}

	// Try year-month format: 2022-04
	if t, err := time.Parse("2006-01", s); err == nil {
		return t.Format("Jan 2006")
	}

	// Return as-is (probably just a year like "2022")
	return s
}

// Template functions available in all templates
var funcMap = template.FuncMap{
	"escape": escapeTypst,
	"date":   formatDate,
}

// unescapeHTML converts HTML entities back to their original characters
func unescapeHTML(s string) string {
	// First use Go's html.UnescapeString for standard entities
	s = html.UnescapeString(s)
	// Handle any remaining numeric entities
	s = strings.ReplaceAll(s, "&#39;", "'")
	s = strings.ReplaceAll(s, "&39;", "'")
	s = strings.ReplaceAll(s, "&#43;", "+")
	return s
}

func main() {
	log.SetFlags(0)
	var infile, outfile, templateName, validatePDF string
	var showText bool

	flag.StringVarP(&infile, "infile", "i", "resume.json", "the file to read in")
	flag.StringVarP(&outfile, "outfile", "o", "", "the file to output to")
	flag.StringVarP(&templateName, "template", "t", "classic", "the template to use")
	flag.StringVarP(&validatePDF, "validate", "v", "", "validate a PDF for ATS compatibility")
	flag.BoolVar(&showText, "show-text", false, "show extracted text when validating")
	flag.Parse()

	// If validate flag is set, run validation and exit
	if validatePDF != "" {
		result, err := validate.ValidatePDF(validatePDF)
		if err != nil {
			log.Fatalf("Validation failed: %v", err)
		}
		log.Print(result.Report())
		if showText {
			log.Print("\n" + strings.Repeat("=", 40) + "\n")
			log.Print("Extracted Text:\n")
			log.Print(result.RawText)
		}
		return
	}

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

	// Normalize legacy field names to v1.0.0 spec
	resume.Normalize()

	// Create template data with both resume and computed fields
	templateData := struct {
		models.Resume
		GroupedWorkEntries []models.GroupedWork
	}{
		Resume:             resume,
		GroupedWorkEntries: resume.GroupedWorkEntries(),
	}

	// Setup the Typst template (using default {{ }} delimiters which don't conflict with Typst's [ ])
	tmpl, err := template.New(infile).Funcs(funcMap).Parse(templateMap[templateName])
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatal(err)
	}

	// Exectute the template
	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, templateData)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatal(err)
	}

	// Unescape HTML entities that may have been introduced
	typstString := unescapeHTML(buf.String())

	// Print out to stdout or file based on flags
	if outfile != "" {
		out, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.SetFlags(log.Lshortfile)
			log.Fatal(err)
		}
		defer out.Close()

		_, err = out.Write([]byte(typstString))
		if err != nil {
			log.SetFlags(log.Lshortfile)
			log.Fatal(err)
		}
	} else {
		log.Print(typstString)
	}
}

var templateMap = map[string]string{
	"classic": templates.Classic,
	"modern":  templates.Modern,
}
