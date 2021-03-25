package utils

import (
	"html/template"
	"net/http"
)


var templates *template.Template

func LoadTemplates(pattern string) {
	// parse the code from the folder templates
	templates = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}