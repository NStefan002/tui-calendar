package google

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

func FetchEvents(srv *calendar.Service, viewing time.Time) (map[string][]*calendar.Event, error) {
	start := time.Date(viewing.Year()-30, 0, 0, 0, 0, 0, 0, viewing.Location())
	end := time.Date(viewing.Year()+30, 0, 0, 0, 0, 0, 0, viewing.Location())

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
		var dateKey string
		if date == "" { // all-day event
			dateKey = item.Start.Date
		} else {
			t, err := time.Parse(time.RFC3339, date)
			if err != nil {
				log.Println("Error parsing date:", err)
				continue
			}
			dateKey = t.Format("2006-01-02")
		}
		events[dateKey] = append(events[dateKey], item)
	}

	return events, nil
}

func CreateEvent(srv *calendar.Service, event *calendar.Event) (*calendar.Event, error) {
	createdEvent, err := srv.Events.Insert("primary", event).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to create event: %w", err)
	}
	return createdEvent, nil
}

func DeleteEvent(srv *calendar.Service, eventID string) error {
	err := srv.Events.Delete("primary", eventID).Do()
	if err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}
	return nil
}
