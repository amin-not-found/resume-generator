package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type ContactInfo struct {
	Display   string
	URL       string
	Icon      string
	IconColor string
}

type Association struct {
	Title       string
	Description string
	Where       string
	Period      string
	URL         string
	Points      []string
	// TODO tags?
}

type Skill struct {
	Name        string
	Description string
}

type SkillSet struct {
	Name   string
	Skills []Skill
}

type Project struct {
	Name        string
	Description string
	URL         string
	Period      string
	Points      []string
	Tags        []string
}

type Summary struct {
	Text   string
	Points []string
}

type Font struct {
	FontFamily string
	FontPath   string
}

type Headers struct {
	Experience     string
	Education      string
	Skills         string
	Projects       string
	Certifications string
}

type PageData struct {
	Name           string
	Summary        Summary
	ContactInfos   []ContactInfo
	Experiences    []Association
	Educations     []Association
	SkillSets      []SkillSet
	Projects       []Project
	Certifications []Association
	Headers        Headers
	Font           Font
	RTL            bool
	// TODO : achievements, activities, etc
}

func exit_if_errf(err error, format string, a ...any) {
	if err != nil {
		fmt.Fprintf(os.Stderr, format, a...)
		fmt.Fprintf(os.Stderr, "\n%v", err)
		os.Exit(1)
	}
}

func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s template_folder input output \n", filepath.Base(args[0]))
		os.Exit(1)
	}
	const base_template = "index.html"
	templates_glob := filepath.Join(args[1], "*.html")
	data_path := filepath.Join(args[2])
	out_path := filepath.Join(args[3])
	// Parse the templates
	templates, err := template.ParseGlob(templates_glob)
	exit_if_errf(err, "Error parsing templates:")

	// Default data
	pageData := PageData{
		Headers: Headers{
			Experience:     "Experience",
			Education:      "Education",
			Skills:         "Skills",
			Projects:       "Projects",
			Certifications: "Certifications",
		},
		Font: Font{"sans-serif", ""},
		RTL:  false,
	}
	// Read data
	data, err := os.ReadFile(data_path)
	exit_if_errf(err, "Error reading from \"%s\":", data_path)
	err = json.Unmarshal(data, &pageData)
	exit_if_errf(err, "Error reading json from \"%s\":", data_path)

	// Create output file
	out_file, err := os.Create(out_path)
	exit_if_errf(err, "Error creating \"%s\":", out_path)

	// Execute templates on output file
	err = templates.ExecuteTemplate(out_file, base_template, pageData)
	exit_if_errf(err, "Error executing templates:")

}
