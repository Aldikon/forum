package handlers

import "html/template"

var templatePage *template.Template

// Fills the templatePage variable with patterns from the path.
func ReadTemplate(pathTemplate string) (err error) {
	templatePage, err = template.ParseGlob(pathTemplate + "*.html")
	return
}
