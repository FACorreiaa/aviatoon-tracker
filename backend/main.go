package main

import (
	"github.com/create-go-app/net_http-go-template/app/routes"
	"github.com/create-go-app/net_http-go-template/docs"
	_ "github.com/create-go-app/net_http-go-template/docs" // load Swagger docs
	"github.com/create-go-app/net_http-go-template/platform/configs"
	"github.com/create-go-app/net_http-go-template/platform/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
	// load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs for Golang net/http Template.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// List of app routes:
	routes.PublicRoutes(r)
	routes.PrivateRoutes(r)
	docs.SwaggerRoutes(r)

	// Register middleware.
	//r.Use(mux.CORSMethodMiddleware(r)) // enable CORS

	// Initialize server.
	server := configs.ServerConfig(r)
	//db, err := database.OpenDBConnection()
	//if err != nil {
	//	panic(err)
	//}
	//err = db.CreateCountryTable()
	//if err != nil {
	//	error.Error(err)
	//}

	//body, err := api.GetAPICountries()
	//fmt.Println(body)
	//if err != nil {
	//	log.Printf("error encoding countries as JSON: %v", err)
	//	return
	//}

	// Start API server.
	utils.StartServerWithGracefulShutdown(server)
}
