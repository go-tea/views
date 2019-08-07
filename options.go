package views

import (
	"fmt"
	"html/template"
	"os"

	"github.com/go-tea/funcs"
)

// Options is a struct for specifying configuration options for the View.Render object.
type Options struct {
	// Asset function to use in place of directory. Defaults to nil.
	Asset func(name string) ([]byte, error)
	// AssetNames function to use in place of directory. Defaults to nil.
	AssetNames func() []string
	//
	AssetInfo func(name string) (os.FileInfo, error)
	//
	AssetDir func(name string) ([]string, error)
	// Directory to load layout templates. Default is "templates/layouts".
	LayoutDir string
	// Directory to load containers templates. Default is "templates/containers".
	ContainersDir string
	// Directory to load includes templates. Default is "templates/includes".
	IncludesDir string
	// Extensions to parse template files from. Defaults to [".tmpl"].
	Extensions []string
	// Funcs is a slice of FuncMaps to apply to the template upon compilation. This is useful for helper functions. Defaults to [].
	Funcs []template.FuncMap
	// Allows changing of output to XHTML instead of HTML. Default is "text/html".
	HTMLContentType string
	// If IsDevelopment is set to true, this will recompile the templates on every request. Default is false.
	//	IsDevelopment bool
	// Appends the given character set to the Content-Type header. Default is "UTF-8".
	Charset string

	// JSONPrefix set Prefix in JSON response
	JSONPrefix string
	// JSONIndent set JSON Indent in response; default false
	JSONIndent bool
	// XMLPrefix set Prefix in XML response
	XMLPrefix string
	// XMLIndent set XML Indent in response; default false
	XMLIndent bool
	// UnEscapeHTML set UnEscapeHTML for JSON; default false
	UnEscapeHTML bool
}

var opt Options

func init() {
	PrepareOptions()
}

func prepareOptions(opts ...Options) Options {
	var opt Options

	if len(opts) > 0 {
		opt = opts[0]
	}

	if len(opt.LayoutDir) == 0 {
		opt.LayoutDir = "templates/layouts"
	}
	if _, err := os.Stat(opt.LayoutDir); os.IsNotExist(err) {
		panic(fmt.Sprintf("LayoutDir %s not exist", opt.LayoutDir))
	}

	if len(opt.ContainersDir) == 0 {
		opt.ContainersDir = "templates/containers"
	}
	if _, err := os.Stat(opt.ContainersDir); os.IsNotExist(err) {
		panic(fmt.Sprintf("Containers %s not exist", opt.ContainersDir))
	}

	if len(opt.IncludesDir) == 0 {
		opt.IncludesDir = "templates/includes"
	}

	if len(opt.Charset) == 0 {
		opt.Charset = "UTF-8"
	}

	if len(opt.Extensions) == 0 {
		opt.Extensions = []string{".tmpl"}
	}

	if len(opt.HTMLContentType) == 0 {
		opt.HTMLContentType = "text/html"
	}

	if opt.Funcs == nil {
		opt.Funcs = []template.FuncMap{funcs.DefaultFuncs()}
	}

	return opt
}

// PrepareOptions function
func PrepareOptions(opts ...Options) {
	opt = prepareOptions(opts...)

	for _, funcs := range opt.Funcs {
		for key, value := range funcs {
			AddFuncMap(key, value)
		}
	}
}
