package models

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"google.golang.org/api/calendar/v3"
)

// enum calendarViewMode represents the different modes of the calendar view
type calendarViewMode int

const (
	calendarView calendarViewMode = iota
	eventDetailsView
	addEventView
	editEventView
)

type model struct {
	cm *calendarModel     // submodel containing calendar data
	dm *eventDetailsModel // submodel for viewing event details
	am *addEventModel     // submodel for adding new events

	calendarViewKeys     calendarViewKeyMap     // key bindings for calendar view
	eventDetailsViewKeys eventDetailsViewKeyMap // key bindings for event details view
	addEventViewKeys     addEventViewKeyMap     // key bindings for add event view

	help help.Model // help view model

	calendarService *calendar.Service            // Google Calendar service for API calls
	events          map[string][]*calendar.Event // key: YYYY-MM-DD, value: list of events for that day

	viewMode     calendarViewMode // current view mode of the calendar
	lastViewMode calendarViewMode // last view mode before switching

	screenWidth  int // width of the terminal screen
	screenHeight int // height of the terminal screen

	loading bool // whether the calendar is currently loading data
	spinner spinner.Model

	errMessage string // error message to display, if any
}

func CreateModel(srv *calendar.Service) model {
	s := spinner.New()
	s.Spinner = spinner.Line

	m := model{
		cm:                   newCM(),
		dm:                   newDM(),
		am:                   newAM(),
		calendarViewKeys:     calendarViewKeys,
		eventDetailsViewKeys: eventDetailsViewKeys,
		addEventViewKeys:     addEventViewKeys,
		help:                 help.New(),
		calendarService:      srv,
		events:               make(map[string][]*calendar.Event),
		viewMode:             calendarView,
		lastViewMode:         calendarView,
		screenWidth:          80,
		screenHeight:         24,
		loading:              true,
		errMessage:           "",
		spinner:              s,
	}

	return m
}
