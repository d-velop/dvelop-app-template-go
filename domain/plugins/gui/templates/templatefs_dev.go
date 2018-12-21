// +build !release

package templates

import (
	"github.com/shurcooL/httpfs/filter"
	"net/http"
	"os"
	"path"
)

// TemplateFileSystem contains the template files and maps to a native filesystem during development
var TemplateFileSystem = filter.Keep(http.Dir("./web/"), func(filepath string, file os.FileInfo) bool {
	if file.IsDir() {
		return true
	}
	if path.Ext(filepath) == ".html" {
		return true
	}
	return false
})

