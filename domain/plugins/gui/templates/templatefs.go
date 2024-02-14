package templates

import (
	"github.com/shurcooL/httpfs/filter"
	_ "github.com/shurcooL/vfsgen"
	"net/http"
	"os"
	"path"
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
