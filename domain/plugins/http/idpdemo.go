package http

import (
	"io"
	"net/http"

	"github.com/d-velop/dvelop-sdk-go/contentnegotiation/mediatype"
	"github.com/d-velop/dvelop-sdk-go/idp"
	"github.com/d-velop/dvelop-sdk-go/log"
)

func HandleIdpDemo(assetBasePath string, renderhtml func(w io.Writer, data interface{}, templatename string) error) http.Handler {
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
			switch negotiatedType.String() {
			case "text/html":
				p, err := idp.PrincipalFromCtx(req.Context())
				if err != nil {
					log.Error(req.Context(), err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
				d := &IdpDemoHtmlDto{Title: "Idp Demo", UserId: p.Id, UserName: p.DisplayName}
				d.AssetBasePath = assetBasePath
				err = renderhtml(w, d, "idpdemo.html")
				if err != nil {
					log.Error(req.Context(), err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			}
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	}
	return http.HandlerFunc(fn)
}

type IdpDemoHtmlDto struct {
	BaseHtmlDto
	UserId   string
	UserName string
	Title    string
}
