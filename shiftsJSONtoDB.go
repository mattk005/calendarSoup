package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"wfm-json-to-ics/internal/database"

	"github.com/google/uuid"
)

func (cfg *config) shiftsJSONtoDB(s *shifts, u *database.User) error {
	dateInFormat := "20060102 150405"
	dateOutFormat := "02 Jan 2006"
	// dateStampFormat := "20060102T150405"
	loc := time.Local // this will probably blow up (change in daylight savings here)
	// time off
	tx, err := cfg.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := cfg.queries.WithTx(tx)
	for _, e := range s.Ltas {
		startTime, err := time.ParseInLocation(dateInFormat, e.StartTime, loc)
		if err != nil {
			fmt.Println(err)
		}
		endTime, err := time.ParseInLocation(dateInFormat, e.EndTime, loc)
		if err != nil {
			fmt.Println(err)
		}
		if err := qtx.AddEvent(context.Background(), database.AddEventParams{
			ID:         uuid.New().String(),
			Summary:    e.TimeOffType,
			EventStart: startTime,
			EventEnd:   endTime,
			Description: sql.NullString{
				String: fmt.Sprintf("Updated: %v", time.Now().Format(dateOutFormat)),
				Valid:  true,
			},
			Location: sql.NullString{},
			UserID:   u.ID,
		}); err != nil {
			fmt.Println(err)
			return err
		}
	}

	// shifts
	for _, e := range s.Shifts {
		startTime, err := time.ParseInLocation(dateInFormat, e.StartTime, loc)
		if err != nil {
			fmt.Println(err)
		}
		endTime, err := time.ParseInLocation(dateInFormat, e.EndTime, loc)
		if err != nil {
			fmt.Println(err)
		}
		if err := qtx.AddEvent(context.Background(), database.AddEventParams{
			ID:         uuid.New().String(),
			Summary:    e.ShiftLabel,
			EventStart: startTime,
			EventEnd:   endTime,
			Description: sql.NullString{
				String: fmt.Sprintf("Updated: %v", time.Now().Format(dateOutFormat)),
				Valid:  true,
			},
			Location: sql.NullString{},
			UserID:   u.ID,
		}); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return tx.Commit()
}
