package http

import (
	"github.com/d-velop/dvelop-sdk-go/contentnegotiation/mediatype"
	"github.com/d-velop/dvelop-sdk-go/log"
	"io"
	"net/http"
	"strings"
)

func HandleNewVacationRequest(assetBasePath string, renderhtml func(w io.Writer, data interface{}, templatename string) error) http.Handler {
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
			d := &VacationRequestDto{Title: "Apply for vacation", Mode: "new"}
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

func HandleVacationRequest(pattern string, assetBasePath string, renderhtml func(w io.Writer, data interface{}, templatename string) error) http.Handler {
	handleList := HandleVacationRequestList(assetBasePath, renderhtml)
	handleSingle := HandleSingleVacationRequest(assetBasePath, renderhtml)
	fn := func(w http.ResponseWriter, req *http.Request) {
		subresource := req.URL.Path[len(pattern):]
		if subresource == "" {
			handleList.ServeHTTP(w, req)
		} else if strings.Contains(subresource, "/") {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleSingle.ServeHTTP(w, req)
		}
	}
	return http.HandlerFunc(fn)
}

func HandleSingleVacationRequest(assetBasePath string, renderhtml func(w io.Writer, data interface{}, templatename string) error) http.Handler {
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
			d := &VacationRequestDto{Title: "Request", Mode: "show"}
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

func HandleVacationRequestList(assetBasePath string, renderhtml func(w io.Writer, data interface{}, templatename string) error) http.Handler {
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
			d := &VacationRequestsDto{Title: "Vacationrequests"}
			d.AssetBasePath = assetBasePath
			err = renderhtml(w, d, "vacationrequests.html")
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

type VacationRequestsDto struct {
	BaseHtmlDto
	Title string
}

type VacationRequestDto struct {
	BaseHtmlDto
	Title string
	Mode  string
}
