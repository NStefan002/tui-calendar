package utils

import (
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"google.golang.org/api/calendar/v3"
)

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

// truncates a string to fit within maxWidth, adding "..." if truncated,
// comparing the width of the styled string to maxWidth
func TruncateString(s string, maxWidth int, style lipgloss.Style) string {
	styled := style.Render(s)
	if lipgloss.Width(styled) <= maxWidth {
		return s
	}

	runes := []rune(s)
	for len(runes) > 0 {
		candidate := style.Render(string(runes) + "...")
		if lipgloss.Width(candidate) <= maxWidth {
			return string(runes) + "..."
		}
		runes = runes[:len(runes)-1]
	}
	return "..."
}

func GetEditor() string {
	editor := "vim" // default editor
	if envEditor := os.Getenv("EDITOR"); envEditor != "" {
		editor = envEditor
	}
	return editor
}
