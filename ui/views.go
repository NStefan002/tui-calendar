package ui

import (
	"fmt"
	"strings"
	"time"

	// "github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func (m model) calendarView() string {
	var sb strings.Builder

	// header (month and year)
	sb.WriteString(headerStyle.Render(m.cm.viewing.Format("January 2006")) + "\n\n\n")

	// days of the week (Mon-Sun)
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for _, day := range daysOfWeek {
		sb.WriteString(baseStyle.Render(fmt.Sprintf("%3s", day)) + " ")
	}
	sb.WriteString("\n\n")

	firstDay := time.Date(m.cm.viewing.Year(), m.cm.viewing.Month(), 1, 0, 0, 0, 0, m.cm.viewing.Location())
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
		isToday := day.Year() == m.cm.now.Year() && day.Month() == m.cm.now.Month() && day.Day() == m.cm.now.Day()
		isSelected := day.Year() == m.cm.selected.Year() && day.Month() == m.cm.selected.Month() && day.Day() == m.cm.selected.Day()

		var dayStr string
		if isSelected {
			dayStr = selectedDateStyle.Render(fmt.Sprintf("%3d", day.Day()))
		} else if isToday {
			dayStr = todayStyle.Render(fmt.Sprintf("%3d", day.Day()))
		} else if m.hasEvents(day) {
			dayStr = dateWithEventStyle.Render(fmt.Sprintf("%3d", day.Day()))
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
	dateKey := m.cm.selected.Format("2006-01-02")
	if events, ok := m.events[dateKey]; ok && len(events) > 0 {
		eventsHeader := eventHeaderStyle.Render("Events for " + m.cm.selected.Format("January 2, 2006"))
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
	return centerText(sb.String(), m.screenWidth)
}

func (m model) eventsView() string {
	dateKey := m.cm.selected.Format("2006-01-02")
	events := m.events[dateKey]
	if len(events) == 0 {
		return centerText("No events for this day.", m.screenWidth)
	}

	selected := events[m.dm.idx]

	// left column: full details
	var details strings.Builder
	details.WriteString(eventStyle.Render(fmt.Sprintf("Title: %s\n", selected.Summary)))
	if selected.Location != "" {
		details.WriteString(eventStyle.Render(fmt.Sprintf("Location: %s\n", selected.Location)))
	}
	if selected.Start != nil {
		details.WriteString(eventStyle.Render(fmt.Sprintf("\nStart: %s\n", formatTime(selected.Start))))
	}
	if selected.End != nil {
		details.WriteString(eventStyle.Render(fmt.Sprintf("\nEnd:   %s\n", formatTime(selected.End))))
	}
	if selected.Description != "" {
		details.WriteString(eventStyle.Render("\nDescription:\n" + selected.Description + "\n"))
	}

	leftCol := eventDetailsStyle.Render(details.String())

	// right column: list of events
	var list strings.Builder
	for i, event := range events {
		timeStr := "All-day"
		if event.Start != nil && event.Start.DateTime != "" {
			t, err := time.Parse(time.RFC3339, event.Start.DateTime)
			if err == nil {
				timeStr = t.Format("15:04")
			}
		}
		line := fmt.Sprintf("%s  %s", timeStr, event.Summary)
		if i == m.dm.idx {
			list.WriteString(eventListSelectedStyle.Render(line) + "\n")
		} else {
			list.WriteString(eventListStyle.Render(line) + "\n")
		}
	}
	rightCol := list.String()

	// compose side by side
	combined := lipgloss.JoinHorizontal(lipgloss.Top, leftCol, rightCol)

	// footer navigation
	footer := "\n[j/k] Move  [escq] Back"
	return centerText(combined+footer, m.screenWidth)
}

func (m model) editEventView() string {
	return "EDIT EVENT VIEW (not implemented yet)"
}

func (m model) addEventView() string {
	return "ADD EVENT VIEW (not implemented yet)"
}
