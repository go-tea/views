package views

import (
	"fmt"
	"html/template"
	"text/template/parse"
)

type Template struct {
	*template.Template
}

func NewTemplate(name string) *Template {
	tmpl := template.New(name)
	return &Template{tmpl}
}

func (t *Template) Parse(text string) *Template {
	tmpl := templateMust(t.Template.Parse(text))
	t.Template = tmpl
	return t
}

func (t *Template) ParseFiles(filenames ...string) *Template {
	tmpl := templateMust(t.Template.ParseFiles(filenames...))
	t.Template = tmpl
	return t
}

func (t *Template) ParseDir(dirpath string) *Template {
	files := filesFromDir(dirpath)
	tmpl := templateMust(t.Template.ParseFiles(files...))
	t.Template = tmpl
	return t
}

func (t *Template) ParseGlob(pattern string) *Template {
	tmpl := templateMust(t.Template.ParseGlob(pattern))
	t.Template = tmpl

	return t
}

func (t *Template) AddParseTree(childname string, tree *parse.Tree) *Template {
	tmpl := templateMust(t.Template.AddParseTree(childname, tree))
	t.Template = tmpl
	return t
}

func (t *Template) Lookup(name string) *template.Template {
	return t.Template.Lookup(name)
}

func (t *Template) AddTemplate() {
	templates[t.Template.Name()] = t.Template
}

func GetTemplate(name string) *Template {
	tmpl, ok := templates[name]

	if !ok {
		panic(fmt.Errorf("The template %s does not exist.\n", name))
	}
	return &Template{tmpl}
}
