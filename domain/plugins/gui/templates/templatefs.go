package templates

import (
	"net/http"
	"os"
	"path"

	"github.com/shurcooL/httpfs/filter"
)

var HTMLTemplates = filter.Keep(http.Dir("./web/"), func(filepath string, file os.FileInfo) bool {
	if file.IsDir() {
		return true
	}
	if path.Ext(filepath) == ".html" {
		return true
	}
	return false
})
