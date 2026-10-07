package getEventsJson

import (
	"context"
	"fmt"

	"wfm-json-to-ics/internal/config"
)

func updateFromWFM(cfg *config.Config) error {
	usersDB, err := cfg.Queries.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, u := range usersDB {

		eventsJSON, err := getEventsFromWfm(u, cfg)
		// fmt.Println(shiftsJSON)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if err = shiftsJSONtoDB(&eventsJSON, &u, cfg); err != nil {
			return err
		}
		// icsCal, err := makeCalendar(&u, cfg) // icsCal is a pointer
		// if err != nil {
		// 	return err
		// }
		//
		// if err := localics.WriteICSOut(icsCal, &u, cfg); err != nil {
		// 	return err
		// }
	}
	return nil
}
