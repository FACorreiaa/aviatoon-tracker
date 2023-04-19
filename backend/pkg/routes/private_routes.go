package routes

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/create-go-app/net_http-go-template/app/controllers"
	"github.com/create-go-app/net_http-go-template/pkg/configs"
	"github.com/go-chi/chi/v5"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(router *chi.Mux) {
	// Define JWT middleware.
	jwtProtected := jwtmiddleware.New(configs.JWTConfig())

	// Define JWT protected routes.
	createUser := jwtProtected.Handler(http.HandlerFunc(controllers.CreateUser))
	updateUser := jwtProtected.Handler(http.HandlerFunc(controllers.UpdateUser))
	deleteUser := jwtProtected.Handler(http.HandlerFunc(controllers.DeleteUser))

	// Routes for POST method:
	router.Post("/api/v1/user", createUser.(http.HandlerFunc))

	// Routes for PUT method:
	router.Put("/api/v1/user", updateUser.(http.HandlerFunc))

	// Routes for DELETE method:
	router.Delete("/api/v1/user", deleteUser.(http.HandlerFunc))
}
