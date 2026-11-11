package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
)

// Group a set of global varable in a struct
var opts struct {
	http struct {
		port int
	}
	logger struct {
		level string
	}
}

type application struct {
	logger    *slog.Logger
	templates map[string]*template.Template
}

var loggerlevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func main() {
	flag.IntVar(&opts.http.port, "http-port", 6767, "The HTTP port to run the server on")
	flag.StringVar(&opts.logger.level, "log-level", "info", "Minimal log level (debug, info, warn, error)")
	flag.Parse()

	logLevel, ok := loggerlevels[opts.logger.level]
	if !ok {
		log.Fatalf("invalid log level supplied: %s", opts.logger.level)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	logger.Info("logger initialized", "level", opts.logger.level)

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:    logger,
		templates: templateCache,
	}

	err = run(app, opts.http.port)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func run(app *application, port int) error {

	app.logger.Info("starting server", "port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		return err
	}

	return nil
}
