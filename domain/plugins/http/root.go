package http

import (
	"encoding/json"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/conf"
	"io"
	"net/http"

	"github.com/d-velop/dvelop-sdk-go/contentnegotiation/mediatype"
	"github.com/d-velop/dvelop-sdk-go/log"
)

func HandleRoot(assetBasePath string, renderhtml func(w io.Writer, data interface{}, templatename string) error, version string) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		// todo 404, if /subressource is requestd
		switch req.Method {
		case http.MethodGet:
			negotiatedType, err := mediatype.Negotiate(req.Header.Get("Accept"), []string{"text/html", "application/hal+json", "application/json"})
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
				return
			}
			w.Header().Set("content-type", negotiatedType.String()+";charset=utf-8")
			switch negotiatedType.String() {
			case "text/html":
				d := &RootHtmlDto{Title: "Vacationprocess", Version: version}
				d.AssetBasePath = assetBasePath
				err = renderhtml(w, d, "root.html")
				if err != nil {
					log.Error(req.Context(), err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			default:
				// application/hal+json or application/json
				err := json.NewEncoder(w).Encode(&HalResource{Links: map[string]HalLinkObject{
					"featuresdescription": {Href: conf.BasePath + "/features"},
				}})
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

type BaseHtmlDto struct {
	AssetBasePath string
}

type RootHtmlDto struct {
	BaseHtmlDto
	Title   string
	Version string
}

// HalResource represents a HAL Resource Object as defined in https://tools.ietf.org/html/draft-kelly-json-hal-08#section-4
type HalResource struct {
	Links map[string]HalLinkObject `json:"_links"`
}

// HalLinkObject represents a HAL Link Object as defined in https://tools.ietf.org/html/draft-kelly-json-hal-08#section-5
type HalLinkObject struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated,omitempty"`
}
