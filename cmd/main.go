package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/julianyf/encurtador-de-url/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed execute code", "error", err)
		return
	}
	slog.Info("All systems offline")
}

func run() error {
	db := make(map[string]string)

	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  10 + time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 + time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
