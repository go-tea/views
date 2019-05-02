package views

import (
	"fmt"
	"html/template"
)

var tplFuncMap = make(template.FuncMap)

// AddFuncMap let user to register a func in the template.
func AddFuncMap(key string, fn interface{}) error {
	tplFuncMap[key] = fn
	return nil
}

func PrintFuncs() {
	for key, value := range tplFuncMap {
		fmt.Printf("%s:\t %T\n", key, value)
	}
}
