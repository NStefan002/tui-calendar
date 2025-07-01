package ui

import (
	"fmt"
	"sort"
	"strings"
	"time"

	// "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"google.golang.org/api/calendar/v3"
)

func CreateModel() model {
	return model{
		now:             time.Now(),
		viewing:         time.Now(),
		selected:        time.Now(),
		calendarService: nil,
		events:          make(map[string][]*calendar.Event),
		viewMode:        CalendarView,
		screenWidth:     80,
		screenHeight:    24,
		loading:         false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.screenWidth = msg.Width
		m.screenHeight = msg.Height

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
		case "r":
			m.loading = true
			// TODO:
			// m.events = fetchEvents(m.calendarService, m.viewing)
			m.loading = false
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
		sb.WriteString("      ")
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

		dateKey := day.Format("2006-01-02")
		if events, ok := m.events[dateKey]; ok && len(events) > 0 {
			sort.Slice(events, func(i, j int) bool {
				startTimeI, _ := time.Parse(time.RFC3339, events[i].Start.DateTime)
				startTimeJ, _ := time.Parse(time.RFC3339, events[j].Start.DateTime)
				return startTimeI.Before(startTimeJ)
			})
			for _, event := range events {
				eventTime, err := time.Parse(time.RFC3339, event.Start.DateTime)
				if err != nil {
					continue
				}
				eventStr := eventStyle.Render(fmt.Sprintf(" %s: %s", eventTime.Format("15:04"), event.Summary))
				sb.WriteString("\n" + eventStr)
			}
		}

		// break line at Sunday (weekday = 0)
		w := int(day.Weekday())
		if w == 0 {
			sb.WriteString("\n\n")
		}
	}

	return sb.String()
}
