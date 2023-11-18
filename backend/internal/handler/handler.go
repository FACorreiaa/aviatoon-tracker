package handler

import (
	"context"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/external_api"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/pprof"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/prometheus"
	"github.com/FACorreiaa/aviatoon-tracker/internal/service"
	"github.com/FACorreiaa/aviatoon-tracker/pkg/logs"
	"os"
	"sync"
	"syscall"
)

type handler interface {
	Run() error
	Shutdown(ctx context.Context) error
}

type Config struct {
	externalApiConfig external_api.Config
	pprofConfig       pprof.Config
	prometheusConfig  prometheus.Config
}

func NewConfig(
	apiConfig external_api.Config,
	pprofConfig pprof.Config,
	prometheusConfig prometheus.Config,
) Config {
	return Config{
		externalApiConfig: apiConfig,
		pprofConfig:       pprofConfig,
		prometheusConfig:  prometheusConfig,
	}
}

type Handler struct {
	service *service.Service
	config  Config
	ctx     context.Context

	externalApi handler
	pprof       handler
	prometheus  handler
}

func NewHandler(
	c Config,
	s *service.Service,
) *Handler {
	return &Handler{
		config:  c,
		service: s,
	}
}

func (h *Handler) Handle(exitSignal *os.Signal) {
	h.externalApi = external_api.New(h.config.externalApiConfig, h.service)
	h.pprof = pprof.New(h.config.pprofConfig)
	h.prometheus = prometheus.New(h.config.prometheusConfig)
	go func() {
		if err := h.pprof.Run(); err != nil && exitSignal == nil {
			logs.DefaultLogger.WithError(err).Fatal("Pprof server was closed unexpectedly")
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		}
	}()
	go func() {
		if err := h.externalApi.Run(); err != nil && exitSignal == nil {
			logs.DefaultLogger.WithError(err).Fatal("REST API Server was closed unexpectedly")
			syscall.Kill(syscall.Getpid(), syscall.SIGQUIT)
		}
	}()
	go func() {
		if err := h.prometheus.Run(); err != nil && exitSignal == nil {
			logs.DefaultLogger.WithError(err).Fatal("Prometheus was closed unexpectedly")
			syscall.Kill(syscall.Getpid(), syscall.SIGQUIT)
		}
	}()
}

func (h *Handler) Shutdown(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		if err := h.externalApi.Shutdown(ctx); err != nil {
			logs.DefaultLogger.WithError(err).Fatal("Error on restApi shutdown")
		}
		wg.Done()
	}()
	go func() {
		if err := h.pprof.Shutdown(ctx); err != nil {
			logs.DefaultLogger.WithError(err).Fatal("Error on pprof shutdown")
		}
		wg.Done()
	}()
	go func() {
		if err := h.prometheus.Shutdown(ctx); err != nil {
			logs.DefaultLogger.WithError(err).Fatal("Error on pprof shutdown")
		}
		wg.Done()
	}()
	wg.Wait()
}
