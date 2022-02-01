package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
)

type page struct {
	TemplateFileName string
	HTMLPagePath     string
	Content          string
}

func main() {
	fileFlag := flag.String("file", "first-post.txt", "// The content the page will have. **MUST ADD .txt TO END of FILE NAME**")

	fmt.Println(*fileFlag)

	dat, _ := os.ReadFile(*fileFlag)
	newPage := page{"template.tmpl", "first-post.html", string(dat)}

	t := template.Must(template.New(newPage.TemplateFileName).ParseFiles(newPage.TemplateFileName))

	newFile, err := os.Create(newPage.HTMLPagePath)
	if err != nil {
		panic(err)
	}

	t.Execute(newFile, newPage)
}
