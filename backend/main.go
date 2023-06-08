package main

import (
	"context"

	"github.com/FACorreiaa/aviatoon-tracker/configs"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/external_api"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/pprof"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/prometheus"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FACorreiaa/aviatoon-tracker/internal/repository"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres"
	"github.com/FACorreiaa/aviatoon-tracker/internal/service"
	"github.com/FACorreiaa/aviatoon-tracker/pkg/logs"
	"github.com/joho/godotenv"
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
			prometheus.NewConfig(
				config.Handlers.Prometheus.Port,
				config.Handlers.Prometheus.KeyFile,
				config.Handlers.Prometheus.CertFile,
				config.Handlers.Prometheus.EnableTLS,
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

// func getHandlerMode(mode string) handler.Mode {
// 	switch mode {
// 	case "prod":
// 		return handler.Production
// 	case "test":
// 		return handler.Test
// 	case "dev":
// 		return handler.Development
// 	default:
// 		logs.DefaultLogger.Fatal("Mode has no match")
// 		return ""
// 	}
// }

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
