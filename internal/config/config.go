package config

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"wfm-json-to-ics/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	BaseURL   string
	Client    *http.Client
	Tokens    Tokens
	DateRange DateRange
	Queries   *database.Queries
	DB        *sql.DB
}

func New() Config {
	dbPath := os.Getenv("DBPATH")
	if dbPath == "" {
		dbPath = "./calendarSoup.db"
	}
	dbConn, err := sql.Open("sqlite3", dbPath+"?_journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}

	cfg := Config{
		BaseURL:   os.Getenv("BASEURL"),
		Client:    &http.Client{},
		DateRange: getDateRange(),
		Queries:   database.New(dbConn),
		DB:        dbConn,
	}
	return cfg
}
