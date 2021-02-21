package views

import "html/template"

// NewView returns the View type that contains template type
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

// View contains template type
type View struct {
	Template *template.Template
}
