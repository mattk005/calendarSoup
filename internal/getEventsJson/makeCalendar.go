package getEventsJson

import (
	"context"

	"wfm-json-to-ics/internal/config"
	"wfm-json-to-ics/internal/database"

	ics "github.com/arran4/golang-ical"
	"github.com/google/uuid"
)

func makeCalendar(u *database.User, cfg *config.Config) (*ics.Calendar, error) {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	cal.SetDescription("My Schedule")
	cal.SetTzid("America/Vancouver")
	cal.SetProductId("I don't want to talk about it")

	// dateOutFormat := "02 Jan 2006"
	dateStampFormat := "20060102T150405"
	// loc := time.Local
	events, err := cfg.Queries.GetEventsByUser(context.Background(), u.ID)
	if err != nil {
		return cal, err
	}
	for _, e := range events {
		event := cal.AddEvent(uuid.New().String())
		event.SetSummary(e.Summary)
		event.SetDescription(e.Description.String)
		event.SetStartAt(e.EventStart)
		event.SetEndAt(e.EventEnd)
		event.SetProperty(ics.ComponentPropertyDtstamp, e.UpdatedAt.Format(dateStampFormat))
	}

	return cal, nil
}
