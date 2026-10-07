package main

import (
	"log"

	"wfm-json-to-ics/internal/config"
	httpserver "wfm-json-to-ics/internal/httpServer"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("couldn't load .env file")
	}

	cfg := config.New()
	defer cfg.DB.Close()
	httpserver.StartServer(&cfg)
}
