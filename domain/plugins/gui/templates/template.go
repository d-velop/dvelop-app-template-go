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
	files, err := getFilesIn(TemplateFileSystem, "/")
	if err != nil {
		log.Error(context.Background(), err)
		panic(err)
	}
	t, e := vfstemplate.ParseFiles(TemplateFileSystem, nil, files...)
	if e != nil {
		log.Error(context.Background(), e)
		panic(e)
	}
	return t
}

// getFilesIn returns all files from the given file system fs starting from the
// given root directory.
func getFilesIn(fs http.FileSystem, root string) ([]string, error) {
	var filenames []string
	d, err := fs.Open(root)
	if err != nil {
		return nil, err
	}
	files, err := d.Readdir(-1)
	if err != nil {
		return nil, err
	}
	for _, fi := range files {
		if fi.IsDir() {
			f, err := getFilesIn(fs, path.Join(root, fi.Name()))
			if err != nil {
				return nil, err
			}
			filenames = append(filenames, f...)
			continue
		}
		filenames = append(filenames, path.Join(root, fi.Name()))
	}
	return filenames, nil
}
