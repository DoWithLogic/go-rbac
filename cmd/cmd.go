package cmd

import (
	"context"
	"strings"
	"time"

	"github.com/DoWithLogic/go-rbac/config"
	"github.com/DoWithLogic/go-rbac/docs"
	"github.com/DoWithLogic/go-rbac/internal/server"
	"github.com/labstack/gommon/log"
)

func Start() {
	// Load the application configuration from the specified directory.
	cfg, err := config.LoadConfig("config")
	if err != nil {
		panic(err)
	}

	// make swagger host dynamic.
	docs.SwaggerInfo.Host = cfg.App.Host
	docs.SwaggerInfo.Schemes = strings.Split(cfg.App.Scheme, ",")
	docs.SwaggerInfo.Version = cfg.App.Version

	if _, err := time.LoadLocation(cfg.Server.TimeZone); err != nil {
		panic(err)
	}

	// Create a new instance of the application and start it.
	if err := server.NewServer(context.Background(), cfg).Run(); err != nil {
		log.Warn(err)
	}
}
