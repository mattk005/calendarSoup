package main

import (
	"context"
	"fmt"
)

func (cfg *config) updateFromWFM() error {
	usersDB, err := cfg.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, u := range usersDB {

		shiftsJSON, err := cfg.getShifts(u)
		fmt.Println(shiftsJSON)
		if err != nil {
			fmt.Println(err)
			return err
		}
		iscCal := makeCalendar(shiftsJSON)
		fmt.Println(iscCal.Serialize())
		break
	}
	return nil
}
