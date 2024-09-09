package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"go-groc-store/config"
	"go-groc-store/pkg/database"
)

type Server struct {
	port int

	db database.Service
}

func NewServer(logger *slog.Logger, cfg *config.Config) *http.Server {
	port, _ := strconv.Atoi(cfg.Http.Port)
	fmt.Println(port)
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
