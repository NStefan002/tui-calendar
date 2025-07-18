package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"google.golang.org/api/calendar/v3"
)

// enum CalendarViewMode represents the different modes of the calendar view
type CalendarViewMode int

const (
	CalendarView CalendarViewMode = iota
	EventsView
	AddEventView
	EditEventView
)

type CalendarModel struct {
	now      time.Time // current time
	viewing  time.Time // month viewed in the calendar
	selected time.Time // selected date in the calendar
}

type DetailsModel struct {
	idx int // index of the currently selected event in the list
}

type AddEventModel struct {
	title       string    // title of the new event
	description string    // description of the new event
	startTime   time.Time // start time of the new event
	endTime     time.Time // end time of the new event
	location    string    // location of the new event
	allDay      bool      // whether the event is an all-day event

	idx int // index of the text input field being edited
}

type EditEventModel struct {
	title       string    // title of the event
	description string    // description of the event
	startTime   time.Time // start time of the event
	endTime     time.Time // end time of the event
	location    string    // location of the event
	allDay      bool      // whether the event is an all-day event

	event *calendar.Event // event being edited
	idx   int             // index of the text input field being edited
}

type model struct {
	cm CalendarModel  // submodel containing calendar data
	dm DetailsModel   // submodel for viewing event details
	am AddEventModel  // submodel for adding new events
	em EditEventModel // submodel for editing existing events

	calendarService *calendar.Service            // Google Calendar service for API calls
	events          map[string][]*calendar.Event // key: YYYY-MM-DD, value: list of events for that day

	viewMode     CalendarViewMode // current view mode of the calendar
	lastViewMode CalendarViewMode // last view mode before switching

	screenWidth  int // width of the terminal screen
	screenHeight int // height of the terminal screen

	loading bool // whether the calendar is currently loading data
	spinner spinner.Model

	errMessage string // error message to display, if any
}

func CreateModel(srv *calendar.Service) model {
	s := spinner.New()
	s.Spinner = spinner.Line
	return model{
		cm: CalendarModel{
			now:      time.Now(),
			viewing:  time.Now(),
			selected: time.Now(),
		},
		dm: DetailsModel{
			idx: 0,
		},
		am: AddEventModel{
			title:       "",
			description: "",
			startTime:   time.Now(),
			endTime:     time.Now().Add(time.Hour),
			location:    "",
			allDay:      false,
			idx:         0,
		},
		em: EditEventModel{
			title:       "",
			description: "",
			startTime:   time.Now(),
			endTime:     time.Now().Add(time.Hour),
			location:    "",
			allDay:      false,
			event:       nil,
			idx:         0,
		},
		calendarService: srv,
		events:          make(map[string][]*calendar.Event),
		viewMode:        CalendarView,
		lastViewMode:    CalendarView,
		screenWidth:     80,
		screenHeight:    24,
		loading:         true,
		errMessage:      "",
		spinner:         s,
	}
}
