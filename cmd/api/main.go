package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "0.1"

type config struct {
	env  string
	port int
}

type application struct {
	cfg    config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "specify server port")
	flag.StringVar(&cfg.env, "env", "dev", "specify operating environment")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
		cfg:    cfg,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)

	srv := &http.Server{
		Addr: fmt.Sprintf("localhost:%d", cfg.port),
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 5 * time.Second,
    IdleTimeout: time.Minute,
    Handler: mux,
    ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

  logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

  err := srv.ListenAndServe()

  logger.Error(err.Error())
  os.Exit(1)
}
