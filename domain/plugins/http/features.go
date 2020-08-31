package http

import (
	"encoding/json"
	"git.d-velop.de/dvelopcloud/approuter-echoapp/domain/plugins/conf"
	"net/http"

	"github.com/d-velop/dvelop-sdk-go/contentnegotiation/mediatype"
	"github.com/d-velop/dvelop-sdk-go/log"
)

// todo add link as soon as the URL to the HomeApp API is stable
func HandleFeatures() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			negotiatedType, err := mediatype.Negotiate(req.Header.Get("Accept"), []string{"text/html", "application/hal+json", "application/json"})
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
				return
			}
			w.Header().Set("content-type", negotiatedType.String()+";charset=utf-8")
			err = json.NewEncoder(w).Encode(&FeaturesDto{Features: []FeatureDto{
				{
					Url:         conf.BasePath + "/vacationrequest/",
					Title:       "Vacation management",
					Subtitle:    "Manage vacation requests",
					IconURI:     conf.AssetBasePath() + "/feature_icon.svg",
					Summary:     "Your vacation requests and the requests of your employees",
					Description: "Apply for vacation and approve the vacation requests of your employees",
					Color:       "#adff2f",
				},
			}})
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

type FeaturesDto struct {
	Features []FeatureDto `json:"features"`
}

type FeatureDto struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	IconURI     string `json:"iconURI"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Color       string `json:"color"`
}
