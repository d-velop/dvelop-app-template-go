package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func NewEchoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		r := HttpRequest{
			Method: req.Method,
			Url:    req.URL.String(),
			Header: map[string][]string{},
		}

		for k, v := range req.Header {
			r.Header[k] = v
		}

		w.Header().Set("content-type", "application/json")
		err := json.NewEncoder(w).Encode(r)

		if err != nil {
			log.Println("failed to encode as json:", err)
		}
	})
}

type HttpRequest struct {
	Method string              `json:"method"`
	Url    string              `json:"url"`
	Header map[string][]string `json:"header"`
}
