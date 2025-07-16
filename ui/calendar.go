package ui

import (
	"context"
	"fmt"
	"time"

	// "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"google.golang.org/api/calendar/v3"
)

type errMsg struct{ error }

type eventsMsg map[string][]*calendar.Event

func (m model) Init() tea.Cmd {
	m.loading = true
	return tea.Batch(
		m.spinner.Tick, // start spinner on app load
		func() tea.Msg {
			events, err := fetchEvents(m.calendarService, m.viewing)
			if err != nil {
				return errMsg{err}
			}
			return eventsMsg(events)
		},
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.screenWidth = msg.Width
		m.screenHeight = msg.Height

	// custom message for events
	case eventsMsg:
		m.events = msg
		m.loading = false

	// custom message for events
	case errMsg:
		m.loading = false
		m.errMessage = msg.Error()

	case tea.KeyMsg:
		switch m.viewMode {
		case CalendarView:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "left", "h":
				m.selected = m.selected.AddDate(0, 0, -1) // go to previous day
			case "right", "l":
				m.selected = m.selected.AddDate(0, 0, 1) // go to next day
			case "up", "k":
				m.selected = m.selected.AddDate(0, 0, -7) // go to previous week
			case "down", "j":
				m.selected = m.selected.AddDate(0, 0, 7) // go to next week
			case "pageup", "pgup", "ctrl+u":
				m.selected = m.selected.AddDate(0, -1, 0) // go to previous month
			case "pagedown", "pgdown", "ctrl+d":
				m.selected = m.selected.AddDate(0, 1, 0) // go to next month
			case "r":
				m.loading = true
				return m, tea.Batch(
					m.spinner.Tick, // start spinner
					func() tea.Msg {
						events, err := fetchEvents(m.calendarService, m.viewing)
						if err != nil {
							return errMsg{err}
						}
						return eventsMsg(events)
					},
				)
			case "enter":
				m.viewMode = DetailsView
			}

		case DetailsView:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.viewMode = CalendarView
			case "down", "j":
				if len(m.events) > 0 && m.selectedEventIdx < len(m.events[m.selected.Format("2006-01-02")])-1 {
					m.selectedEventIdx++
				}
			case "up", "k":
				if m.selectedEventIdx > 0 {
					m.selectedEventIdx--
				}
			}
		}

	// spinner update
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	if m.selected.Month() != m.viewing.Month() || m.selected.Year() != m.viewing.Year() {
		m.viewing = m.selected // update viewing month if selected date is not in current viewing month
	}

	return m, nil
}

func (m model) View() string {
	if m.loading {
		return centerText(m.spinner.View()+" Loading calendar events...", m.screenWidth)
	}

	if len(m.errMessage) > 0 {
		return "There was an error: " + m.errMessage
	}

	switch m.viewMode {
	case CalendarView:
		return m.calendarView()
	case DetailsView:
		return m.detailsView()
	default:
		return "ERROR"
	}
}

func fetchEvents(srv *calendar.Service, viewing time.Time) (map[string][]*calendar.Event, error) {
	start := time.Date(viewing.Year(), viewing.Month(), 1, 0, 0, 0, 0, viewing.Location())
	end := start.AddDate(0, 1, 0)

	events := make(map[string][]*calendar.Event)

	call := srv.Events.List("primary").
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(start.Format(time.RFC3339)).
		TimeMax(end.Format(time.RFC3339)).
		OrderBy("startTime")

	resp, err := call.Context(context.Background()).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %w", err)
	}

	for _, item := range resp.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date // all-day event
		}
		t, err := time.Parse(time.RFC3339, date)
		if err != nil {
			continue
		}
		dateKey := t.Format("2006-01-02")
		events[dateKey] = append(events[dateKey], item)
	}

	return events, nil
}
