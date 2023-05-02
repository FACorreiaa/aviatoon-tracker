package docs

import (
	"github.com/create-go-app/net_http-go-template/platform/utils"
	"github.com/go-chi/chi/v5"

	httpSwagger "github.com/swaggo/http-swagger"
)

// SwaggerRoutes func for describe group of Swagger routes.
func SwaggerRoutes(router *chi.Mux) {
	// Define server settings:
	serverConnURL, _ := utils.ConnectionURLBuilder("server")

	// Build Swagger route.
	getSwagger := httpSwagger.Handler(
		httpSwagger.URL("http://"+serverConnURL+"/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	)

	// Routes for GET method:
	router.Route("/", func(r chi.Router) {
		r.Get("/swagger/", getSwagger)
	})
}
