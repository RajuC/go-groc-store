package main

import (
	"go-groc-store/config"
	"go-groc-store/pkg/database"
	"go-groc-store/pkg/log"
	"go-groc-store/pkg/server"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger := log.NewLoggerService()
	cfg, er := config.NewConfigService(logger, "config/")
	if er != nil {
		panic(er)
	}
	logger = log.SetLoggerLevel(cfg.Log.Level)

	logger.Info("Confguration", slog.Any("cfg", cfg))
	dbService := database.New(logger, cfg)
	app := fiber.New()

	server := server.NewServer(app, logger, cfg.Http.Port, *dbService)
	server.Start()

}
