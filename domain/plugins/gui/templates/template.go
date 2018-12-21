package templates

import (
	"context"
	"github.com/d-velop/dvelop-sdk-go/log"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	"html/template"
)

func parseTemplates() *template.Template {
	t, e := vfstemplate.ParseGlob(TemplateFileSystem, nil, "*")
	if e != nil {
		log.Error(context.Background(), e)
		panic(e)
	}
	return t
}
