package views

import (
	"html/template"
	"path/filepath"
)

var (
	layoutDir   string = "views/layouts/"
	templateExt string = ".gohtml"
)

// NewView returns the View type that contains template type
func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/footer.gohtml",
		"views/layouts/navbar.gohtml",
		"views/layouts/bootstrap.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View contains template type
type View struct {
	Template *template.Template
	Layout   string
}

// layoutFiles returns a slice of strings representing
// the layout files used in our application
func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files
}
