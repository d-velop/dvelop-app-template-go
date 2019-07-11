package templates

import (
	"context"
	"html/template"
	"net/http"
	"path"

	"github.com/d-velop/dvelop-sdk-go/log"
	"github.com/shurcooL/httpfs/html/vfstemplate"
)

func parseTemplates() *template.Template {
	files, err := listFiles(TemplateFileSystem)
	if err != nil {
		panic(err)
	}
	t, e := vfstemplate.ParseFiles(TemplateFileSystem, nil, files...)
	if e != nil {
		log.Error(context.Background(), e)
		panic(e)
	}
	return t
}

func listFiles(fs http.FileSystem) ([]string, error) {
	var res []string
	if err := readDirRecursive(fs, "/", &res); err != nil {
		return nil, err
	}
	return res, nil
}

func readDirRecursive(fs http.FileSystem, dir string, res *[]string) error {
	f, err := fs.Open(dir)
	if err != nil {
		return err
	}
	files, err := f.Readdir(-1)
	if err != nil {
		return err
	}
	for _, fi := range files {
		if fi.IsDir() {
			if err := readDirRecursive(fs, path.Join(dir, fi.Name()), res); err != nil {
				return err
			}
			continue
		}
		*res = append(*res, path.Join(dir, fi.Name()))
	}
	return nil
}
