// +build ignore

// Command to generate an in memory version of the AssetFileSystem
// cf. https://github.com/shurcooL/vfsgen
// It can be invoked by running go generate
package main

import (
	"flag"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/gui/assets"
	"log"
	"os"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var wd string
	flag.StringVar(&wd, "workdir", "./", "Workingdir. Must be the project root.")
	flag.Parse()
	os.Chdir(wd)

	err := vfsgen.Generate(assets.AssetFileSystem, vfsgen.Options{
		Filename:        "./domain/plugins/gui/assets/assetfs_prod_gen.go",
		PackageName:     "assets",
		BuildTags:       "release",
		VariableName:    "AssetFileSystem",
		VariableComment: "AssetFileSystem contains the asset files in memory for production releases.",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
