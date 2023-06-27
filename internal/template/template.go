package template

import (
	"html/template"
	"os"
	p "path"
	"strings"
)

var Page = make(map[string]*template.Template, 0)

func InitTemplate(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return err
		}

		pathTemplate := p.Join(path, info.Name())
		ts, err := template.ParseFiles(pathTemplate)
		if err != nil {
			return err
		}
		Page[strings.TrimSuffix(info.Name(), ".html")] = ts

	}

	return nil
}
