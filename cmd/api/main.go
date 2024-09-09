package main

import (
	"fmt"
	"go-groc-store/config"
	"go-groc-store/pkg/server"
)

func main() {

	cfg, er := config.NewConfigService()
	if er != nil {
		panic(er)
	}
	fmt.Println(cfg)
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
