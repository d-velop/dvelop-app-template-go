package http

import (
	"net/http"
)

func HandleAssets(pref string, fs http.FileSystem) http.Handler {
	return http.StripPrefix(pref, http.FileServer(fs))
}
