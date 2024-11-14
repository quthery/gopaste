package server

import (
	"gopaste/internal/config"
	"net/http"
	"time"
)

func Run(cfg *config.Config, handler http.Handler) error {
	httpserver := &http.Server{
		Addr:           cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return httpserver.ListenAndServe()
}
