package utils

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"google.golang.org/api/calendar/v3"
)

// center each line of text based on the screen width
func CenterText(text string, width int) string {
	// calculate the padding needed to center the text
	padding := max((width-lipgloss.Width(text))/2, 0)

	// split the text into lines and center each line
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.Repeat(" ", padding) + line
	}
	return strings.Join(lines, "\n")
}

func FormatTime(dt *calendar.EventDateTime) string {
	if dt.DateTime != "" {
		t, err := time.Parse(time.RFC3339, dt.DateTime)
		if err == nil {
			return t.Format("Mon Jan 2, 15:04")
		}
	}
	if dt.Date != "" {
		t, err := time.Parse("2006-01-02", dt.Date)
		if err == nil {
			return t.Format("Mon Jan 2 (All-day)")
		}
	}
	return "Unknown"
}

func HasEvents(events map[string][]*calendar.Event, day time.Time) bool {
	dateKey := day.Format("2006-01-02")
	_, exists := events[dateKey]
	return exists && len(events[dateKey]) > 0
}
