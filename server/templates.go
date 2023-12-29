package server

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
)

func CreateMultiRenderer(templatesDir string) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	configureMultiTemplates(renderer, templatesDir)
	configurePartials(renderer, templatesDir)

	return renderer
}

func configureMultiTemplates(renderer multitemplate.Renderer, templatesDir string) {
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/**/*.tmpl.html")
	if err != nil {
		panic(err.Error())
	}

	printSeparator("=")
	printTemplateInfo("TEMPLATE NAME", "TEMPLATE FILE")
	printSeparator("-")

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)

		templateFileName := strings.Replace(include, "templates/includes/", "", 1)
		templateName := strings.Replace(strings.Replace(templateFileName, ".tmpl.html", "", -1), "/", "_", -1)
		printTemplateInfo(templateName, templateFileName)

		renderer.AddFromFiles(templateName, files...)
	}
}

func configurePartials(renderer multitemplate.Renderer, templatesDir string) {
	partials, err := filepath.Glob(templatesDir + "/partials/**/*.tmpl.html")
	if err != nil {
		panic(err.Error())
	}

	printSeparator("=")
	printTemplateInfo("PARTIAL NAME", "PARTIAL FILE")
	printSeparator("-")

	for _, partial := range partials {
		templateFileName := strings.Replace(partial, "templates/partials/", "", 1)
		templateName := strings.Replace(strings.Replace(templateFileName, ".tmpl.html", "", -1), "/", "_", -1)
		printTemplateInfo(templateName, templateFileName)
		renderer.AddFromFiles(templateName, partial)
	}

	printSeparator("=")
}

func printTemplateInfo(templateName string, templateFileName string) {
	log.Printf("| %-25s | %s", templateName, templateFileName)
}

func printSeparator(separator string) {
	log.Printf("|%s ", strings.Repeat(separator, 80))
}
