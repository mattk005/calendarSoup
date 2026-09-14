package main

import (
	"context"
	"fmt"
)

func (cfg *config) updateFromWFM() error {
	usersDB, err := cfg.queries.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, u := range usersDB {

		shiftsJSON, err := cfg.getShifts(u)
		// fmt.Println(shiftsJSON)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if err = cfg.shiftsJSONtoDB(&shiftsJSON, &u); err != nil {
			return err
		}
		icsCal, err := cfg.makeCalendar(&u) // icsCal is a pointer
		if err != nil {
			return err
		}
		// fmt.Println(iscCal.Serialize())
		cfg.writeICSOut(icsCal, &u)
	}
	return nil
}
