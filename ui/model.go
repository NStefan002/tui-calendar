package ui

import (
	"time"

	"google.golang.org/api/calendar/v3"
)

// enum CalendarViewMode represents the different modes of the calendar view
type CalendarViewMode int

const (
	CalendarView CalendarViewMode = iota
	DetailsView
	AddEventView
	EditEventView
)

type model struct {
	now      time.Time // current time
	viewing  time.Time // month viewed in the calendar
	selected time.Time // selected date in the calendar

	calendarService *calendar.Service // Google Calendar service for API calls

	events map[string][]*calendar.Event // key: YYYY-MM-DD, value: list of events for that day

	viewMode CalendarViewMode // current view mode of the calendar

	screenWidth  int // width of the terminal screen
	screenHeight int // height of the terminal screen

	loading bool // whether the calendar is currently loading data
}

func CreateModel(srv *calendar.Service) model {
	return model{
		now:             time.Now(),
		viewing:         time.Now(),
		selected:        time.Now(),
		calendarService: srv,
		events:          make(map[string][]*calendar.Event),
		viewMode:        CalendarView,
		screenWidth:     80,
		screenHeight:    24,
		loading:         false,
	}
}
