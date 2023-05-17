package prometheus

import (
	"context"
	"syscall"

	"github.com/FACorreiaa/aviatoon-tracker/pkg/logs"
	"github.com/go-chi/chi/v5"
)

type Config struct {
	port      string
	keyFile   string
	certFile  string
	enableTls bool
}

func NewConfig(
	port string,
	keyFile string,
	certFile string,
	enableTls bool,
) Config {
	if enableTls && (certFile == "" || keyFile == "") {
		logs.DefaultLogger.Fatal("Tls is enabled but cert file or key file doesn't have a path")
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}
	return Config{
		port:      port,
		keyFile:   keyFile,
		certFile:  certFile,
		enableTls: enableTls,
	}
}

type Prometheus interface {
	Run() error
	Shutdown(ctx context.Context) error
}

func New(config Config) Prometheus {

	handler := chi.NewMux()
	InitPrometheus(handler)
	return &server{
		port:      config.port,
		handler:   handler,
		certFile:  config.certFile,
		keyFile:   config.keyFile,
		enableTls: config.enableTls,
	}
}
