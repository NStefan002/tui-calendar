package ui

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"google.golang.org/api/calendar/v3"
)

// center each line of text based on the screen width
func centerText(text string, width int) string {
	// calculate the padding needed to center the text
	padding := max((width-lipgloss.Width(text))/2, 0)

	// split the text into lines and center each line
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.Repeat(" ", padding) + line
	}
	return strings.Join(lines, "\n")
}

func formatTime(dt *calendar.EventDateTime) string {
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

func (m model) hasEvents(day time.Time) bool {
	dateKey := day.Format("2006-01-02")
	_, exists := m.events[dateKey]
	return exists && len(m.events[dateKey]) > 0
}

func (am *AddEventModel) nextField() {
	if am.title.Focused() {
		am.title.Blur()
		am.title.PromptStyle = inactiveTextinputStyle
		am.title.TextStyle = inactiveTextinputStyle
		am.title.Cursor.Style = inactiveTextinputStyle

		am.location.Focus()
		am.location.PromptStyle = activeTextinputStyle
		am.location.TextStyle = activeTextinputStyle
		am.location.Cursor.Style = activeTextinputStyle
	} else if am.location.Focused() {
		am.location.Blur()
		am.location.PromptStyle = inactiveTextinputStyle
		am.location.TextStyle = inactiveTextinputStyle
		am.location.Cursor.Style = inactiveTextinputStyle

		am.title.Focus()
		am.title.PromptStyle = activeTextinputStyle
		am.title.TextStyle = activeTextinputStyle
		am.title.Cursor.Style = activeTextinputStyle
	}
}

func (am *AddEventModel) prevField() {
	if am.title.Focused() {
		am.title.Blur()
		am.title.PromptStyle = inactiveTextinputStyle
		am.title.TextStyle = inactiveTextinputStyle
		am.title.Cursor.Style = inactiveTextinputStyle

		am.location.Focus()
		am.location.PromptStyle = activeTextinputStyle
		am.location.TextStyle = activeTextinputStyle
		am.location.Cursor.Style = activeTextinputStyle
	} else if am.location.Focused() {
		am.location.Blur()
		am.location.PromptStyle = inactiveTextinputStyle
		am.location.TextStyle = inactiveTextinputStyle
		am.location.Cursor.Style = inactiveTextinputStyle

		am.title.Focus()
		am.title.PromptStyle = activeTextinputStyle
		am.title.TextStyle = activeTextinputStyle
		am.title.Cursor.Style = activeTextinputStyle
	}
}
