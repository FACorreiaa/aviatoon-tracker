package pprof

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitPprof(router chi.Router) {

	router.Route("/debug/pprof", func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Get("/", chiWrapper(pprof.Index))
		r.Get("/cmdline", chiWrapper(pprof.Cmdline))
		r.Get("/profile", chiWrapper(pprof.Profile))
		r.Post("/symbol", chiWrapper(pprof.Symbol))
		r.Get("/symbol", chiWrapper(pprof.Symbol))
		r.Get("/trace", chiWrapper(pprof.Trace))
		r.Get("/allocs", chiWrapper(handlerFunc(pprof.Handler("allocs"))))
		r.Get("/block", chiWrapper(handlerFunc(pprof.Handler("block"))))
		r.Get("/goroutine", chiWrapper(handlerFunc(pprof.Handler("goroutine"))))
		r.Get("/heap", chiWrapper(handlerFunc(pprof.Handler("heap"))))
		r.Get("/mutex", chiWrapper(handlerFunc(pprof.Handler("mutex"))))
		r.Get("/threadcreate", chiWrapper(handlerFunc(pprof.Handler("threadcreate"))))
	})
}

func handlerFunc(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func chiWrapper(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fn(w, r.WithContext(ctx))
	}
}
