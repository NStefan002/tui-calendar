package ui

import (
	"fmt"
	"strings"
	"time"

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

	// align calendar to start on Monday (make Sunday = 7)
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

	// left column: event titles
	var list strings.Builder
	for i, event := range events {
		title := event.Summary
		if title == "" {
			title = "[No Title]"
		}
		if i == m.dm.idx {
			list.WriteString(eventListSelectedStyle.Render(title) + "\n")
		} else {
			list.WriteString(eventListStyle.Render(title) + "\n")
		}
	}
	leftCol := boxStyle.Width(30).Render(list.String())

	// right column: selected event details
	var right strings.Builder

	// title
	eventTitle := selected.Summary
	if eventTitle == "" {
		eventTitle = "[No Title]"
	}
	right.WriteString(centerText(detailTitleStyle.Render(eventTitle), 50) + "\n\n")

	// times
	var startStr, endStr string
	if selected.Start != nil && selected.Start.DateTime != "" {
		startTime, err := time.Parse(time.RFC3339, selected.Start.DateTime)
		if err == nil {
			startStr = timeLabelStyle.Render("Start: ") + timeValueStyle.Render(startTime.Format("Mon, Jan 2 — 15:04"))
		}
	}
	if selected.End != nil && selected.End.DateTime != "" {
		endTime, err := time.Parse(time.RFC3339, selected.End.DateTime)
		if err == nil {
			endStr = timeLabelStyle.Render("End:   ") + timeValueStyle.Render(endTime.Format("Mon, Jan 2 — 15:04"))
		}
	}
	if startStr != "" {
		right.WriteString(startStr + "\n")
	}
	if endStr != "" {
		right.WriteString(endStr + "\n")
	}

	// description
	desc := strings.TrimSpace(selected.Description)
	if desc == "" {
		desc = "[No description]"
	}
	right.WriteString("\n" + descriptionStyle.Render(desc))

	rightCol := boxStyle.Width(50).Render(right.String())

	// side-by-side layout
	main := lipgloss.JoinHorizontal(lipgloss.Top, leftCol, rightCol)

	// footer
	footer := formFooterStyle.Render("[j/k] Navigate  [q/esc] Back")

	return centerText(main, m.screenWidth) + "\n\n" + centerText(footer, m.screenWidth)
}

func (m model) editEventView() string {
	return "EDIT EVENT VIEW (not implemented yet)"
}

func (m model) addEventView() string {
	var sb strings.Builder

	// header
	header := headerStyle.Render(fmt.Sprintf("➕ Add Event for %s", m.cm.selected.Format("January 2, 2006")))
	sb.WriteString(centerText(header, m.screenWidth) + "\n\n")

	// form fields
	fields := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, fieldLabelStyle.Render("Title:"), m.am.title.View()),
		lipgloss.JoinHorizontal(lipgloss.Top, fieldLabelStyle.Render("Location:"), m.am.location.View()),
		// you can add more like description, start time, end time similarly.
	}

	form := lipgloss.JoinVertical(lipgloss.Left, fields...)
	box := boxStyle.Render(form)

	sb.WriteString(centerText(box, m.screenWidth))

	// footer
	footer := formFooterStyle.Render("[tab] Next field  [enter] Confirm  [esc/q] Cancel")
	sb.WriteString("\n\n" + centerText(footer, m.screenWidth))

	return sb.String()
}
