// +build !release

package assets

import (
	"github.com/shurcooL/httpfs/filter"
	"net/http"
	"os"
	"path"
)

// AssetFileSystem contains the asset files and maps to a native filesystem during development
var AssetFileSystem = filter.Keep(http.Dir("./web/"), func(filepath string, file os.FileInfo) bool {
	if file.IsDir() {
		return true
	}
	if path.Ext(filepath) != ".html" {
		return true
	}
	return false
})
