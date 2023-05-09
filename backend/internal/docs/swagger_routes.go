package docs

import (
	"github.com/go-chi/chi/v5"

	httpSwagger "github.com/swaggo/http-swagger"
)

// SwaggerRoutes func for describe group of Swagger routes.
func SwaggerRoutes(router *chi.Mux) {
	// Define server settings:

	//Swagger
	router.Mount("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("./doc.json"), // Use a relative URL for the Swagger documentation file
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

}
