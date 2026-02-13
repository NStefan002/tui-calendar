package utils

import (
	"time"

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
