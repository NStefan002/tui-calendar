package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/NStefan002/tui-calendar/v2/styles"
	"github.com/NStefan002/tui-calendar/v2/utils"

	"github.com/charmbracelet/lipgloss"
)

type calendarModel struct {
	now      time.Time // current time
	viewing  time.Time // month viewed in the calendar
	selected time.Time // selected date in the calendar
}

func newCM() *calendarModel {
	return &calendarModel{
		now:      time.Now(),
		viewing:  time.Now(),
		selected: time.Now(),
	}
}

func (cm *calendarModel) view(m *model) string {
	var sb strings.Builder

	// header (month and year)
	sb.WriteString(styles.Header.Render(cm.viewing.Format("January 2006")) + "\n\n\n")

	// days of the week (Mon-Sun)
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for _, day := range daysOfWeek {
		sb.WriteString(styles.Base.Render(fmt.Sprintf("%3s", day)) + " ")
	}
	sb.WriteString("\n\n")

	firstDay := time.Date(cm.viewing.Year(), cm.viewing.Month(), 1, 0, 0, 0, 0, cm.viewing.Location())
	lastDay := firstDay.AddDate(0, 1, -1)

	// align calendar to start on Monday (make Sunday = 7)
	weekday := int(firstDay.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	for i := 1; i < weekday; i++ {
		sb.WriteString(strings.Repeat(" ", lipgloss.Width(styles.Base.Render(daysOfWeek[0]))+1))
	}

	for day := firstDay; !day.After(lastDay); day = day.AddDate(0, 0, 1) {
		isToday := day.Year() == cm.now.Year() && day.Month() == cm.now.Month() && day.Day() == cm.now.Day()
		isSelected := day.Year() == cm.selected.Year() && day.Month() == cm.selected.Month() && day.Day() == cm.selected.Day()

		var dayStr string
		if isSelected {
			dayStr = styles.SelectedDate.Render(fmt.Sprintf("%3d", day.Day()))
		} else if isToday {
			dayStr = styles.Today.Render(fmt.Sprintf("%3d", day.Day()))
		} else if utils.HasEvents(m.events, day) {
			dayStr = styles.DateWithEvent.Render(fmt.Sprintf("%3d", day.Day()))
		} else {
			dayStr = styles.Base.Render(fmt.Sprintf("%3d", day.Day()))
		}
		sb.WriteString(dayStr + " ")

		// break line at Sunday (weekday = 0)
		w := int(day.Weekday())
		if w == 0 {
			sb.WriteString("\n\n")
		}
	}

	// display events (if any) for the selected date
	dateKey := cm.selected.Format("2006-01-02")
	if events, ok := m.events[dateKey]; ok && len(events) > 0 {
		eventsHeader := styles.EventHeader.Render("Events for " + cm.selected.Format("January 2, 2006"))
		sb.WriteString("\n\n\n" + eventsHeader + "\n")
		for _, event := range events {
			eventTimeStr := ""
			if event.Start.DateTime == "" {
				// All-day event
				eventTimeStr = styles.TimeValue.Render("All Day")
			} else {
				eventTime, err := time.Parse(time.RFC3339, event.Start.DateTime)
				if err != nil {
					log.Println("Error parsing event time:", err)
					continue
				}
				eventTimeStr = styles.TimeValue.Render(eventTime.Format("15:04"))
			}
			titleStr := ""
			if event.Summary == "" {
				titleStr = "(No Title)"
			} else {
				titleStr = event.Summary
			}
			eventTitle := styles.Event.Render(titleStr)
			eventTimeTitleGap := strings.Repeat(" ", lipgloss.Width(eventsHeader)-lipgloss.Width(eventTimeStr)-lipgloss.Width(eventTitle))
			sb.WriteString(fmt.Sprintf("\n%s%s%s", eventTimeStr, eventTimeTitleGap, eventTitle))
		}
	}

	sb.WriteString("\n")

	main := utils.CenterText(sb.String(), m.screenWidth)
	helpText := utils.CenterText(m.help.View(m.calendarViewKeys), m.screenWidth)

	return main + "\n\n" + helpText
}
