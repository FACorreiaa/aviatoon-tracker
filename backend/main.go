package main

import (
	"context"
	"github.com/FACorreiaa/aviatoon-tracker/configs"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/external_api"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/pprof"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres"
	"github.com/FACorreiaa/aviatoon-tracker/internal/service"
	"github.com/FACorreiaa/aviatoon-tracker/pkg/logs"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logs.InitDefaultLogger()
	config, err := configs.InitConfig()
	if err != nil {
		logs.DefaultLogger.WithError(err).Error("Config was not configure")
	}
	logs.DefaultLogger.Info("Config was successfully imported")
	logs.DefaultLogger.ConfigureLogger(
		getLogFormatter(config.Mode),
	)
	logs.DefaultLogger.Info("Main logger was initialized successfully")

	if err := godotenv.Load(config.Dotenv); err != nil && config.Dotenv != "" {
		logs.DefaultLogger.WithError(err).Fatal("Dotenv was not loaded")
		os.Exit(1)
	}
	logs.DefaultLogger.Info("Dotenv file was successfully loaded")

	repositories := repository.NewRepository(
		repository.NewConfig(
			postgres.NewConfig(
				config.Repositories.Postgres.Host,
				config.Repositories.Postgres.Port,
				config.Repositories.Postgres.Username,
				os.Getenv("POSTGRES_PASSWORD"),
				config.Repositories.Postgres.DB,
				config.Repositories.Postgres.SSLMode,
				10*time.Second,
				postgres.CacheStatement,
			),
		),
	)
	logs.DefaultLogger.Info("Repository was initialized")
	services := service.NewService(repositories)
	logs.DefaultLogger.Info("Service was initialized")
	handlers := handler.NewHandler(
		handler.NewConfig(
			external_api.NewConfig(
				config.Handlers.ExternalApi.Port,
				config.Handlers.ExternalApi.KeyFile,
				config.Handlers.ExternalApi.CertFile,
				config.Handlers.ExternalApi.EnableTLS,
			),
			pprof.NewConfig(
				config.Handlers.Pprof.Port,
				config.Handlers.Pprof.KeyFile,
				config.Handlers.Pprof.CertFile,
				config.Handlers.Pprof.EnableTLS,
			),
		),
		services,
	)
	logs.DefaultLogger.Info("Handler was initialized")

	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT,
	)
	var exitSignal os.Signal
	handlers.Handle(&exitSignal)
	logs.DefaultLogger.Info("Handler was successfully started")
	exitSignal = <-quit
	logs.DefaultLogger.Info("Exit...")
	handlers.Shutdown(context.Background())
	logs.DefaultLogger.Info("Handlers are shutdown")
}

//func getHandlerMode(mode string) handler.Mode {
//	switch mode {
//	case "prod":
//		return handler.Production
//	case "test":
//		return handler.Test
//	case "dev":
//		return handler.Development
//	default:
//		logs.DefaultLogger.Fatal("Mode has no match")
//		return ""
//	}
//}

func getLogFormatter(mode string) logs.Formatter {
	switch mode {
	case "prod":
		return logs.JSONFormatter
	case "test":
		return logs.DefaultFormatter
	case "dev":
		return logs.DefaultFormatter
	default:
		logs.DefaultLogger.Fatal("Mode has no match")
		os.Exit(1)
		return 0
	}
}

//package main
//
//import (
//	"github.com/FACorreiaa/aviatoon-tracker/docs"
//	_ "github.com/FACorreiaa/aviatoon-tracker/docs" // load Swagger docs
//	"github.com/FACorreiaa/aviatoon-tracker/pkg/configs"
//	"github.com/FACorreiaa/aviatoon-tracker/pkg/routes"
//	"github.com/FACorreiaa/aviatoon-tracker/pkg/utils"
//	"github.com/go-chi/chi/v5"
//	"github.com/go-chi/chi/v5/middleware"
//	_ "github.com/joho/godotenv/autoload"
//	// load .env file automatically
//)
//
//// @title API
//// @version 1.0
//// @description This is an auto-generated API Docs for Golang net/http Template.
//// @termsOfService http://swagger.io/terms/
//// @contact.name API Support
//// @contact.email your@mail.com
//// @license.name Apache 2.0
//// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//// @securityDefinitions.apikey ApiKeyAuth
//// @in header
//// @name Authorization
//// @BasePath /api
//func main() {
//	r := chi.NewRouter()
//	r.Use(middleware.Logger)
//
//	// List of app routes:
//	routes.PublicRoutes(r)
//	routes.PrivateRoutes(r)
//	docs.SwaggerRoutes(r)
//
//	// Register middleware.
//r.Use(mux.CORSMethodMiddleware(r)) // enable CORS
//
//	// Initialize server.
//	server := configs.ServerConfig(r)
//	//db, err := database.OpenDBConnection()
//	//if err != nil {
//	//	panic(err)
//	//}
//	//err = db.CreateCountryTable()
//	//if err != nil {
//	//	error.Error(err)
//	//}
//
//	//body, err := api.GetAPICountries()
//	//fmt.Println(body)
//	//if err != nil {
//	//	log.Printf("error encoding countries as JSON: %v", err)
//	//	return
//	//}
//
//	// Start API server.
//	utils.StartServerWithGracefulShutdown(server)
//}
