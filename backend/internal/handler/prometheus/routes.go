package prometheus

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func InitPrometheus(router chi.Router) {

	router.Route("/prometheus", func(r chi.Router) {
		r.Use(middleware.NoCache)
		//r.Get("/metrics", chiWrapper(pprof.Index))
		//r.Get("/cmdline", chiWrapper(pprof.Cmdline))
		//r.Get("/profile", chiWrapper(pprof.Profile))
		//r.Post("/symbol", chiWrapper(pprof.Symbol))
		//r.Get("/symbol", chiWrapper(pprof.Symbol))
		//r.Get("/trace", chiWrapper(pprof.Trace))
		r.Get("/metrics", chiWrapper(handlerFunc(promhttp.Handler())))

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
