package http

import (
	"io"
	"net/http"

	"github.com/d-velop/dvelop-sdk-go/contentnegotiation/mediatype"
	"github.com/d-velop/dvelop-sdk-go/log"
)

func HandleVacationRequest(assetBasePath string, renderhtml func(w io.Writer, data interface{}, templatename string) error) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			negotiatedType, err := mediatype.Negotiate(req.Header.Get("Accept"), []string{"text/html"})
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
				return
			}
			w.Header().Set("content-type", negotiatedType.String()+";charset=utf-8")
			d := &VacationRequestDto{Title: "Vacationrequest"}
			d.AssetBasePath = assetBasePath
			err = renderhtml(w, d, "vacationrequest.html")
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return

			}
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	}
	return http.HandlerFunc(fn)
}

type VacationRequestDto struct {
	BaseHtmlDto
	Title string
}
