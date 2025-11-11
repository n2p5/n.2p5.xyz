package main

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
)

type HomeData struct {
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	MetaDescription string   `json:"meta_description"`
	Items           []string `json:"items"`
}

type MediaData struct {
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	MetaDescription string      `json:"meta_description"`
	Items           []MediaItem `json:"items"`
}

type MediaItem struct {
	Title   string `json:"title"`
	Details string `json:"details"`
}

func main() {
	if err := generate("data/home.json", "dst/index.html", "home.gohtml", &HomeData{}); err != nil {
		log.Fatal(err)
	}
	if err := generate("data/media.json", "dst/writing-speaking-and-press.html", "media.gohtml", &MediaData{}); err != nil {
		log.Fatal(err)
	}
}

// generate loads JSON data, parses a template, and writes the result to an HTML file.
func generate[T any](inputPath, outputPath, templatePath string, data T) error {
	// Load JSON data
	inFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	if err := json.NewDecoder(inFile).Decode(&data); err != nil {
		return err
	}


	// Parse templates with custom function for raw HTML
	funcMap := template.FuncMap{
		"raw": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	// Parse only the templates needed for this page to avoid name conflicts
	tmpl, err := template.New("").Funcs(funcMap).ParseFiles(
		"templates/style.css",
		"templates/page.gohtml",
		"templates/"+templatePath)
	if err != nil {
		return err
	}

	// Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Execute template
	return tmpl.ExecuteTemplate(outFile, templatePath, data)
}
