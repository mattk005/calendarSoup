package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"wfm-json-to-ics/internal/database"
)

// TODO add cashing

func (cfg *config) getShifts(u database.User) (shifts, error) {
	var s shifts
	if !u.WfmID.Valid {
		return s, fmt.Errorf("this user has no wfmID")
	}
	url, err := url.Parse(cfg.baseURL)
	if err != nil {
		return s, err
	}
	fullURL := url.JoinPath(u.WfmID.String, "shifts")
	q := fullURL.Query()
	q.Add("from", cfg.dateRange.startDate)
	q.Add("to", cfg.dateRange.endDate)
	fullURL.RawQuery = strings.ReplaceAll(q.Encode(), "+", "%20")
	// fmt.Println(fullURL.String())
	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		return s, err
	}
	req.AddCookie(&http.Cookie{
		Name:  "JSESSIONID",
		Value: cfg.tokens.jSessionID,
	})

	req.AddCookie(&http.Cookie{
		Name:  "ROUTEID",
		Value: cfg.tokens.routeID,
	})

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := cfg.client.Do(req)
	if err != nil {
		return s, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return s, fmt.Errorf("bad Status :%s", resp.Status)
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&s)
	if err != nil {
		println(err)
		return s, err
	}
	return s, nil
}
