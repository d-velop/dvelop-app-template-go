// +build !release

package templates

import (
	"io"
)

func Render (w io.Writer, data interface{}, templatename string) error{
	t := parseTemplates() // don't cache parsed templates during development to get changed html templates without the need to restart the process
	return t.ExecuteTemplate(w,templatename,data)
}
