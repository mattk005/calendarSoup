package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"wfm-json-to-ics/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type config struct {
	baseURL   string
	client    *http.Client
	tokens    tokens
	dateRange dateRange
	db        *database.Queries
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("couldn't load .env file")
	}
	dbConn, err := sql.Open("sqlite3", "./calendarSoup.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	token, err := getTokens(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	queries := database.New(dbConn)
	cfg := config{
		baseURL:   os.Getenv("BASEURL"),
		client:    &http.Client{},
		tokens:    token,
		dateRange: getDateRange(),
		db:        queries,
	}

	if err := cfg.getUserData("users.json"); err != nil {
		log.Fatal(err)
	}
	if err := cfg.updateFromWFM(); err != nil {
		log.Fatal(err)
	}
}
