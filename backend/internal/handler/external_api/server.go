package external_api

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

type server struct {
	httpServer *http.Server
	port       string
	handler    http.Handler
	certFile   string
	keyFile    string
	enableTls  bool
}

func (s *server) Run() error {
	var tlsConfig tls.Config
	s.httpServer = &http.Server{
		Addr:           ":" + s.port,
		Handler:        s.handler,
		TLSConfig:      &tlsConfig,
		ReadTimeout:    time.Duration(10) * time.Second,
		WriteTimeout:   time.Duration(10) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server starting on port%+v\n", s.httpServer.Addr)

	//fmt.Printf("%+v\n", s.httpServer.Addr)
	//fmt.Printf("%+v\n", s.httpServer.Handler)
	//fmt.Printf("%+v\n", s.httpServer.TLSConfig)
	//fmt.Printf("%+v\n", s.httpServer.ReadTimeout)
	//fmt.Printf("%+v\n", s.httpServer.WriteTimeout)
	//fmt.Printf("%+v\n", s.httpServer.MaxHeaderBytes)

	if s.enableTls {
		crt, _ := tls.LoadX509KeyPair(s.certFile, s.keyFile)
		tlsConfig = tls.Config{Certificates: []tls.Certificate{crt}}
		return s.httpServer.ListenAndServeTLS("", "")
	}
	return s.httpServer.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
