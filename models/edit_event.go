package models

import (
	"time"

	"google.golang.org/api/calendar/v3"
)

type editEventModel struct {
	title       string    // title of the event
	description string    // description of the event
	startTime   time.Time // start time of the event
	endTime     time.Time // end time of the event
	location    string    // location of the event
	allDay      bool      // whether the event is an all-day event

	event *calendar.Event // event being edited
	idx   int             // index of the text input field being edited
}

func newEM() *editEventModel {
	return &editEventModel{
		title:       "",
		description: "",
		startTime:   time.Now(),
		endTime:     time.Now().Add(time.Hour),
		location:    "",
		allDay:      false,
		event:       nil,
		idx:         0,
	}
}

func (em *editEventModel) view() string {
	return "EDIT EVENT VIEW (not implemented yet)"
}
