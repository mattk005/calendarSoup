package main

import (
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/google/uuid"
)

func makeCalendar(shifts shifts) *ics.Calendar {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	cal.SetDescription("My Schedule")
	cal.SetTzid("America/Vancouver")
	cal.SetProductId("I don't want to talk about it")

	dateInFormat := "20060102 150405"
	dateOutFormat := "02 Jan 2006"
	dateStampFormat := "20060102T150405"
	loc := time.Local

	// time off
	for _, s := range shifts.Ltas {
		event := cal.AddEvent(uuid.NewString())
		event.SetSummary(s.TimeOffType)

		descriptionString := fmt.Sprintf("Updated: %v", time.Now().Format(dateOutFormat))
		event.SetDescription(descriptionString)

		startTime, err := time.ParseInLocation(dateInFormat, s.StartTime, loc)
		if err != nil {
			fmt.Println(err)
		}
		event.SetStartAt(startTime)
		endTime, err := time.ParseInLocation(dateInFormat, s.EndTime, loc)
		if err != nil {
			fmt.Println(err)
		}
		event.SetEndAt(endTime)
		// TODO whoops, hard coded datestamp (coppied from gemini)
		event.SetProperty(ics.ComponentPropertyDtstamp, time.Now().Format(dateStampFormat))
	}

	// shifts
	for _, s := range shifts.Shifts {
		event := cal.AddEvent(uuid.NewString())
		event.SetSummary(s.ShiftLabel)
		if len(s.Details) > 0 {
			event.SetLocation(s.Details[0].Department.LocDesc)
		}

		descriptionString := fmt.Sprintf("Updated: %v", time.Now().Format(dateOutFormat))
		event.SetDescription(descriptionString)
		startTime, err := time.ParseInLocation(dateInFormat, s.StartTime, loc)
		if err != nil {
			fmt.Println(err)
		}
		event.SetStartAt(startTime)
		endTime, err := time.ParseInLocation(dateInFormat, s.EndTime, loc)
		if err != nil {
			fmt.Println(err)
		}
		event.SetEndAt(endTime)
		event.SetProperty(ics.ComponentPropertyDtstamp, time.Now().Format(dateStampFormat))
	}

	return cal
}
