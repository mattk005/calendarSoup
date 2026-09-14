package main

import (
	"database/sql"
	"fmt"
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
	queries   *database.Queries
	db        *sql.DB
}

func main() {
	// // 1. Create CPU profile file
	// f, err := os.Create("cpu.pprof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	//
	// // 2. Start CPU profiling
	// if err := pprof.StartCPUProfile(f); err != nil {
	// 	log.Fatal(err)
	// }
	// defer pprof.StopCPUProfile()

	args := os.Args
	if len(args) == 1 {
		fmt.Println("this app requires arguments:")
		fmt.Println("\t-updateUsers")
		fmt.Println("\t\tWill update the database users table from users.json")
		fmt.Println("")
		fmt.Println("otherwise, this app requires a JSESSIONID and ROUTEID token from the app.")
		os.Exit(0)
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal("couldn't load .env file")
	}
	dbConn, err := sql.Open("sqlite3", "./calendarSoup.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	queries := database.New(dbConn)
	cfg := config{
		baseURL:   os.Getenv("BASEURL"),
		client:    &http.Client{},
		dateRange: getDateRange(),
		queries:   queries,
		db:        dbConn,
	}

	switch args[1] {
	case "-updateUsers":
		{
			if err = cfg.getUserData("users.json"); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
	case "-printUsers":
		{
			if err = cfg.printUsers(); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}

	}

	cfg.tokens, err = getTokens(args)
	if err != nil {
		log.Fatal(err)
	}

	if err := cfg.updateFromWFM(); err != nil {
		log.Fatal(err)
	}
	// // 3. Optional: Capture Memory allocation profile at the end
	// memFile, err := os.Create("mem.pprof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer memFile.Close()
	// pprof.WriteHeapProfile(memFile)
}
