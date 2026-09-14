package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"wfm-json-to-ics/internal/database"
)

type tokens struct {
	jSessionID string
	routeID    string
}

type userData struct {
	User string `json:"user"`
	ID   string `json:"id"`
	Num  string `json:"num"`
}

func tokenHelper(s, sep string) string {
	_, rest, ok := strings.Cut(s, sep)
	if !ok {
		log.Fatalf("failed to find %s in token string", sep)
	}
	result, _, ok := strings.Cut(rest, ";")
	if !ok {
		log.Fatalf("failed to find %s in token string", sep)
	}
	return result
}

func getTokens(args []string) (tokens, error) {
	argsString := strings.Join(args, "") + ";"
	t := tokens{
		jSessionID: tokenHelper(argsString, "JSESSIONID="),
		routeID:    tokenHelper(argsString, "ROUTEID="),
	}
	return t, nil
}

func (cfg *config) getUserData(path string) error {
	fp, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	file, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	var users []userData
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&users); err != nil {
		return err
	}
	for _, u := range users {
		_, err = cfg.queries.CreateUser(context.Background(), database.CreateUserParams{
			Name: u.User,
			WfmID: sql.NullString{
				String: u.Num,
				Valid:  len(u.Num) > 0,
			},
			WfmExternalCalendar: u.ID,
		})
		if err != nil && !strings.Contains(err.Error(), "UNIQUE constraint failed: users.wfm_external_calendar") {
			return err
		}
	}
	return nil
}
