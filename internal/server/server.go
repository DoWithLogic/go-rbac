package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DoWithLogic/go-rbac/config"
	"github.com/DoWithLogic/go-rbac/pkg/datasources"
	appEcho "github.com/DoWithLogic/go-rbac/pkg/echo"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type Server struct {
	cfg       config.Config
	db        *gorm.DB
	echo      *echo.Echo
	redisConn *redis.Client
}

func NewServer(ctx context.Context, cfg config.Config) *Server {
	// Set a timeout for the database initialization process to avoid indefinite waiting.
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	return &Server{
		cfg:       cfg,
		db:        datasources.NewMySQLDB(ctx, cfg.Database),
		echo:      appEcho.NewEchoServer(cfg.Server),
		redisConn: datasources.NewRedisClient(ctx, cfg.Redis),
	}
}

func (s *Server) Run() error {
	if err := s.mapHandler(); err != nil {
		return err
	}

	// Set up signal handling to gracefully shutdown the server upon receiving a SIGTERM or SIGINT signal.
	// Using a buffered channel with capacity 1 to ensure signals are not missed.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)

	// Start a goroutine to handle the shutdown signal.
	go func() {
		// Wait for the signal from the quit channel.
		<-quit

		// Log the shutdown process.
		log.Info("Server is shutting down...")

		// Create a context with a timeout of 10 seconds for the server shutdown.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		db, err := s.db.DB()
		if err != nil {
			log.Error("failed to retrieve database instance")
		}

		if db != nil {
			if err := db.Close(); err != nil {
				log.Error("failed to close database gracefully")
			}
		}

		if err := s.echo.Shutdown(ctx); err != nil {
			log.Error("failed to shutdown echo gracefully")
		}

		if err := s.redisConn.Close(); err != nil {
			log.Error("failed to close redis gracefully")
		}
	}()

	// Start the echo server and listen on the configured port.
	return s.echo.Start(fmt.Sprintf(":%s", s.cfg.Server.Port))
}
