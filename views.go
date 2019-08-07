package views

//https://elithrar.github.io/article/approximating-html-template-inheritance/

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/oxtoacart/bpool"
)

var bufpool *bpool.BufferPool

func init() {
	bufpool = bpool.NewBufferPool(64)
}

type View struct {
	template *template.Template
	Layout   string
	Vars     map[string]interface{}
}

func GetView(layout string, name string) *View {

	tmpl, ok := templates[name]

	if !ok {
		panic(fmt.Errorf("The template %s does not exist.\n", name))
	}

	return &View{
		template: tmpl,
		Layout:   layout,
		Vars:     make(map[string]interface{}),
	}
}

func (v *View) AddTemplates(fn ...string) {
	ctmpl, _ := templates[v.template.Name()].Clone()

	delete(templates, v.template.Name())
	fid := filesFromIncludeDir(fn...)
	atmpl := template.Must(ctmpl.ParseFiles(fid...))
	templates[v.template.Name()] = atmpl
	v.template = templates[v.template.Name()]

}

func (v *View) AddFuncs(fc ...template.FuncMap) {

	ctmpl, _ := templates[v.template.Name()].Clone()
	delete(templates, v.template.Name())

	var addFuncMap = make(template.FuncMap)

	for _, fs := range fc {
		for key, value := range fs {
			addFuncMap[key] = value
		}
	}

	atmpl := ctmpl.Funcs(addFuncMap)
	templates[v.template.Name()] = atmpl
	v.template = templates[v.template.Name()]

}

func (v *View) Render(w http.ResponseWriter) error {

	// Create a buffer to temporarily write to and check if any errors were encountered.
	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := v.template.ExecuteTemplate(buf, v.Layout, v.Vars)
	if err != nil {
		return StatusError{http.StatusInternalServerError, err}
	}

	// Set the header and write the buffer to the http.ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)
	return nil
}

////// ???  //////////////////////////////////////////////////////////////////////////////////////
func CreateViews(layout []string) map[string]*View {
	v := make(map[string]*View)
	for _, ly := range layout {
		for _, tpl := range templates {
			name := strings.TrimSuffix(tpl.Name(), filepath.Ext(tpl.Name()))
			v[ly+"_"+name] = NewView(ly, tpl.Name())
		}
	}

	return v
}

func NewView(layout string, name string) *View {

	tmpl, ok := templates[name]

	if !ok {
		panic(fmt.Errorf("The template %s does not exist.\n", name))
	}

	return &View{
		template: tmpl,
		Layout:   layout,
		Vars:     make(map[string]interface{}),
	}
}
