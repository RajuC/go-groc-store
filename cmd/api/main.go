package main

import (
	"fmt"
	"go-groc-store/config"
	"go-groc-store/pkg/log"
	"go-groc-store/pkg/server"
)

func main() {
	logger := log.NewLoggerService()
	cfg, er := config.NewConfigService(logger, "config/")
	if er != nil {
		panic(er)
	}
	logger = log.SetLoggerLevel(cfg.Log.Level)
	server := server.NewServer(logger, cfg)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
