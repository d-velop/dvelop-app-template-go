package http

import (
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/conf"
	"github.com/d-velop/dvelop-sdk-go/log"
	"github.com/d-velop/dvelop-sdk-go/requestid"
	"github.com/d-velop/dvelop-sdk-go/requestlog"
	"github.com/d-velop/dvelop-sdk-go/tenant"
	chain "github.com/justinas/alice"
	"net"
	"net/http"

	"context"
)

// Chain which is identical for every resource
var commonHandlerChain = chain.New(
	tenant.AddToCtx(conf.DefaultSystemBaseURI(), conf.SecretKey()),
	requestid.AddToCtx(),
	requestlog.Log(func(ctx context.Context, logmessage string) { log.Info(ctx, logmessage) }),
	addVaryResponseHeader(),
)

func addVaryResponseHeader() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			// Vary Header determines which additional header fields should be used
			// to decide if a request can be answered from a cache
			// cf. https://tools.ietf.org/html/rfc7234#section-4.1
			// accept is added because most resources deliver JSON and HTML from the same URI
			// x-dv-sig-1 ist added because most of the responses are tenant specific
			rw.Header().Set("vary", "accept, x-dv-sig-1")
			next.ServeHTTP(rw, req)
		})
	}
}

func Serve(l net.Listener, handler http.Handler) error {
	return http.Serve(l, handler)
}

// Handle first invokes the common handler chain and then corresponding resource
func Handle(resources []Resource) http.Handler {
	mux := http.NewServeMux()
	for _, r := range resources {
		mux.Handle(r.Pattern, r.Handler)
	}

	return commonHandlerChain.Then(mux)
}

type Resource struct {
	Pattern string
	Handler http.Handler
}
