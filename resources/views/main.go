package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "resources/views/layouts"
var PartialDir string = "resources/views/partials"
var prefix string = ".tmpl"
var baseDir string = "resources/views/"
var bootstrap *template.Template

func NewView(layout string, files ...string) *View {

	for i, filename := range files {
		files[i] = baseDir + filename + prefix
	}

	files = append(layoutFiles(), files...)
	files = append(partialFiles(), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

type ViewData struct {
	Auth interface{}
	Data interface{}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	vd := ViewData{
		Auth: nil,
		Data:    data,
	}
	return v.Template.ExecuteTemplate(w, v.Layout, vd)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*" + prefix)
	if err != nil {
		panic(err)
	}
	return files
}

func partialFiles() []string {
	files, err := filepath.Glob(PartialDir + "/*" + prefix)
	if err != nil {
		panic(err)
	}
	return files
}
