package getEventsJson

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"wfm-json-to-ics/internal/config"
	"wfm-json-to-ics/internal/database"
)

// TODO add cashing

func getEventsFromWfm(u database.User, cfg *config.Config) (shifts, error) {
	var s shifts
	if !u.WfmID.Valid {
		return s, fmt.Errorf("this user has no wfmID")
	}
	url, err := url.Parse(cfg.BaseURL)
	if err != nil {
		return s, err
	}
	fullURL := url.JoinPath("services", "portal", "employee", u.WfmID.String, "shifts")
	q := fullURL.Query()
	q.Add("from", cfg.DateRange.StartDate)
	q.Add("to", cfg.DateRange.EndDate)
	fullURL.RawQuery = strings.ReplaceAll(q.Encode(), "+", "%20")
	// fmt.Println(fullURL.String())
	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		return s, err
	}
	req.AddCookie(&http.Cookie{
		Name:  "JSESSIONID",
		Value: cfg.Tokens.JSessionID,
	})

	req.AddCookie(&http.Cookie{
		Name:  "ROUTEID",
		Value: cfg.Tokens.RouteID,
	})

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := cfg.Client.Do(req)
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
