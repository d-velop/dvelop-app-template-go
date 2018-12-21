// +build release

//go:generate go run templatefs_generate.go --workdir ../../../../

package templates

import (
	"io"
)

var t = parseTemplates() // cache parsed templates for production deployments
func Render (w io.Writer, data interface{}, templatename string) error{
	return t.ExecuteTemplate(w,templatename,data)
}
