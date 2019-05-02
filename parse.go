package views

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
}

func ParseFiles() {

	if opt.Asset == nil || opt.AssetNames == nil {
		compileTemplatesFromDir()
		return
	}
	//todo
	//compileTemplatesFromAsset()

}

func compileTemplatesFromDir() {

	layouts := filesFromDir(opt.LayoutDir)
	//fmt.Println(layouts)
	if layouts == nil {
		panic(fmt.Sprintf("LayoutDir %s is empty", opt.LayoutDir))
	}
	var layoutnames []string
	for _, layout := range layouts {
		templatename := strings.TrimPrefix(layout, opt.LayoutDir+"/")
		templatename = strings.TrimSuffix(templatename, ".tmpl")
		fmt.Println(templatename)
		layoutnames = append(layoutnames, templatename)
	}

	containers := filesFromDir(opt.ContainersDir)
	if containers == nil {
		panic(fmt.Sprintf("ContainersDir %s is empty", opt.ContainersDir))
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, container := range containers {
		files := append(layouts, container)
		templatename := strings.TrimPrefix(filepath.ToSlash(container), opt.ContainersDir+"/")
		err := parseTmpl(templatename, files...)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func parseTmpl(templatename string, files ...string) (e error) {

	var err error
	templates[templatename], err = template.New(templatename).Funcs(tplFuncMap).ParseFiles(files...)
	return err
}

//TODO
func compileTemplatesFromAsset() {
	layouts, err := opt.AssetDir(opt.LayoutDir)
	if err != nil {
		panic(fmt.Sprintf("AssetDir %s err %s", opt.LayoutDir, err))
	}
	fmt.Println(layouts)
	if layouts == nil {
		panic(fmt.Sprintf("LayoutDir %s is empty", opt.LayoutDir))
	}

	containers, err := opt.AssetDir(opt.ContainersDir)
	if err != nil {
		panic(fmt.Sprintf("AssetDir %s err %s", opt.ContainersDir, err))
	}
	fmt.Println(containers)
	if containers == nil {
		panic(fmt.Sprintf("ContainersDir %s is empty", opt.ContainersDir))
	}

}

var files []string

func filesFromDir(dir string) []string {
	files = nil
	filepath.Walk(dir, filepath.WalkFunc(walkpath))
	return files
}

func filesFromIncludeDir(fn ...string) []string {
	var rf []string
	rf = make([]string, 0)
	files := filesFromDir(opt.IncludesDir)
	for _, f := range files {
		for _, n := range fn {
			if n == strings.TrimPrefix(filepath.ToSlash(f), opt.IncludesDir+"/") {
				rf = append(rf, f)
			}
		}
	}
	return rf
}

func walkpath(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if info == nil || info.IsDir() {
		return nil
	}

	ext := filepath.Ext(path)
	for _, extension := range opt.Extensions {
		if ext == extension {
			files = append(files, path)
		}
	}

	return nil
}

func PrintTemplates() {
	for _, tpl := range templates {
		if tpl != nil {
			fmt.Println(tpl.Name())
			fmt.Println(tpl.DefinedTemplates())
		}
	}
}

func templateMust(t *template.Template, err error) *template.Template {
	if err != nil {
		panic("template: " + err.Error())
	}
	return t
}
