// +build ignore

// Command to generate an in memory version of the Templatefilesystem
// cf. https://github.com/shurcooL/vfsgen
// It can be invoked by running go generate
package main

import (
	"flag"
	"git.d-velop.de/dvelopcloud/shop-middleware/domain/plugins/gui/templates"
	"log"
	"os"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var wd string
	flag.StringVar(&wd, "workdir", "./", "Workingdir. Must be the project root.")
	flag.Parse()
	os.Chdir(wd)

	err := vfsgen.Generate(templates.TemplateFileSystem, vfsgen.Options{
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
