package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"wfm-json-to-ics/internal/database"

	ics "github.com/arran4/golang-ical"
)

// actually I problably wont need this
func (cfg *config) writeICSOut(c *ics.Calendar, u *database.User) error {
	fields := strings.Fields(u.Name)
	dir := strings.ToUpper(string(fields[1][0])) + strings.ToUpper(string(fields[0][0]))
	dirPath := filepath.Join("data", dir)
	if err := os.MkdirAll(dirPath, 0o755); err != nil {
		return err
	}
	fp := filepath.Join(dirPath, "calendar.ics")
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	bw := bufio.NewWriter(f)
	if err = c.SerializeTo(bw); err != nil {
		return err
	}
	if err = bw.Flush(); err != nil {
		return err
	}
	return err
}
