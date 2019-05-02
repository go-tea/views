# Views

Parse html all tempates from directory (default templates/containers) to

var templates map[string]*template.Template

~~~ go
// views/views.go

package views

import (
	"flag"
	"html/template"

	"github.com/go-tea/funcs"
	"github.com/go-tea/views"
)

var (
	Home   *views.View
	Bike   *views.View
	Login  *views.View
	Create *views.View
	Show   *views.View
)

var (
	Opts views.Options
)

var (
	servePort = flag.Int("port", 5566, "Port to serve from")
)

func init() {

	
	funcs.SetupLocale("cs_CZ", "messages", "locale")
	funcs.SetupLocale("en_US", "messages", "locale")
	funcs.ChangeLocale("cs_CZ")

	Opts = views.Options{}
	Opts.Funcs = []template.FuncMap{funcs.LangFuncs(), funcs.FormFuncs()}

	views.PrepareOptions(Opts)
	views.ParseFiles()

	Home = views.GetView("main", "home.tmpl")
	Bike = views.GetView("main", "bike.tmpl")
	Bike.AddTemplates("actions.tmpl", "categories.tmpl")
	Create = views.GetView("main", "user/create.tmpl")
	Create.AddFuncs(funcs.FormFuncs())
	Create.AddTemplates("user/form.tmpl")
	Login = views.GetView("main", "user/login.tmpl")
	Login.AddTemplates("user/field.tmpl")
	Show = views.GetView("main", "user/show.tmpl")
	Show.AddFuncs(funcs.FormFuncs())
	Show.AddTemplates("user/user.tmpl")

}

~~~ 