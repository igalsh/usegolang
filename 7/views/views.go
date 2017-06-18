package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "views/layouts"

func NewView(layout string, files ...string) *View {
	// fmt.Fprintf(os.Stdout, files[0])
	files = append(files, layoutFiles()...)
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

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	// for _, fl := range files {
	// 	fmt.Fprintf(os.Stdout, fl)
	// }
	return files
}

func (v *View) Render(w http.ResponseWriter,
	data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprintf(os.Stdout, "entering rendering")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
