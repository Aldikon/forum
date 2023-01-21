package handlers

import "html/template"

var TemplatePage *template.Template

// Fills the templatePage variable with patterns from the path.
func ReadTemplate(pathTemplate string) (err error) {
	TemplatePage, err = template.ParseGlob(pathTemplate + "*.html")
	return
}
