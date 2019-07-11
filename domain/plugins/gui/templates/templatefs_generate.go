// +build ignore

// Command to generate an in memory version of the Templatefilesystem
// cf. https://github.com/shurcooL/vfsgen
// It can be invoked by running go generate
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/vfsgen"
)

func main() {
	var wd string
	flag.StringVar(&wd, "workdir", "./", "Workingdir. Must be the project root.")
	flag.Parse()
	os.Chdir(wd)

	var templateFileSystem = filter.Keep(http.Dir("./web/"), func(filepath string, file os.FileInfo) bool {
		if file.IsDir() {
			return true
		}
		if path.Ext(filepath) == ".html" {
			return true
		}
		return false
	})

	err := vfsgen.Generate(templateFileSystem, vfsgen.Options{
		Filename:        "./domain/plugins/gui/templates/templatefs_prod_gen.go",
		PackageName:     "templates",
		BuildTags:       "release",
		VariableName:    "TemplateFileSystem",
		VariableComment: "TemplateFileSystem contains the template files in memory for production releases.",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
