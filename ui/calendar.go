package ui

import (
	"context"
	"fmt"
	"strings"
	"time"

	// "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"google.golang.org/api/calendar/v3"
)

type errMsg struct{ error }

type eventsMsg map[string][]*calendar.Event

func (m model) Init() tea.Cmd {
	m.loading = true
	return func() tea.Msg {
		events, err := fetchEvents(m.calendarService, m.viewing)
		if err != nil {
			return errMsg{err}
		}
		return eventsMsg(events)
	}
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
		// TODO: optionally set error in model

	case tea.KeyMsg:
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
			return m, func() tea.Msg {
				events, err := fetchEvents(m.calendarService, m.viewing)
				if err != nil {
					return errMsg{err}
				}
				return eventsMsg(events)
			}
		}
	}
	if m.selected.Month() != m.viewing.Month() || m.selected.Year() != m.viewing.Year() {
		m.viewing = m.selected // update viewing month if selected date is not in current viewing month
	}

	return m, nil
}

func (m model) View() string {
	if m.loading {
		return "Loading calendar events..."
	}

	var sb strings.Builder

	// header (month and year)
	sb.WriteString(headerStyle.Render(m.viewing.Format("January 2006")) + "\n\n\n")

	// days of the week (Mon-Sun)
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for _, day := range daysOfWeek {
		sb.WriteString(baseStyle.Render(fmt.Sprintf("%3s", day)) + " ")
	}
	sb.WriteString("\n\n")

	firstDay := time.Date(m.viewing.Year(), m.viewing.Month(), 1, 0, 0, 0, 0, m.viewing.Location())
	lastDay := firstDay.AddDate(0, 1, -1)

	// Align calendar to start on Monday (make Sunday = 7)
	weekday := int(firstDay.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	for i := 1; i < weekday; i++ {
		sb.WriteString(strings.Repeat(" ", lipgloss.Width(baseStyle.Render(daysOfWeek[0]))+1))
	}

	for day := firstDay; !day.After(lastDay); day = day.AddDate(0, 0, 1) {
		isToday := day.Year() == m.now.Year() && day.Month() == m.now.Month() && day.Day() == m.now.Day()
		isSelected := day.Year() == m.selected.Year() && day.Month() == m.selected.Month() && day.Day() == m.selected.Day()

		var dayStr string
		if isSelected {
			dayStr = selectedDateStyle.Render(fmt.Sprintf("%3d", day.Day()))
		} else if isToday {
			dayStr = todayStyle.Render(fmt.Sprintf("%3d", day.Day()))
		} else {
			dayStr = baseStyle.Render(fmt.Sprintf("%3d", day.Day()))
		}
		sb.WriteString(dayStr + " ")

		// break line at Sunday (weekday = 0)
		w := int(day.Weekday())
		if w == 0 {
			sb.WriteString("\n\n")
		}
	}

	// display events (if any) for the selected date
	dateKey := m.selected.Format("2006-01-02")
	if events, ok := m.events[dateKey]; ok && len(events) > 0 {
		eventsHeader := eventHeaderStyle.Render("Events for " + m.selected.Format("January 2, 2006"))
		sb.WriteString("\n\n\n" + eventsHeader + "\n")
		for _, event := range events {
			eventTime, err := time.Parse(time.RFC3339, event.Start.DateTime)
			if err != nil {
				continue
			}
			eventTimeStr := eventStyle.Render(eventTime.Format("15:04"))
			eventTitle := eventStyle.Render(event.Summary)
			eventTimeTitleGap := strings.Repeat(" ", lipgloss.Width(eventsHeader)-lipgloss.Width(eventTimeStr)-lipgloss.Width(eventTitle))
			sb.WriteString(fmt.Sprintf("\n%s%s%s", eventTimeStr, eventTimeTitleGap, eventTitle))
		}
	}

	sb.WriteString("\n")
	return sb.String()
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
