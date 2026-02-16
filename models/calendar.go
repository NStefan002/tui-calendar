package models

import (
	"fmt"
	// "log"
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
	// header (month and year)
	header := styles.Header.Render(cm.viewing.Format("January 2006"))

	// days of the week (Mon-Sun)
	daysOfWeek := func() []string {
		days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
		var result []string
		for i, day := range days {
			result = append(result, styles.Base.Render(day))
			if i < len(days)-1 {
				result = append(result, " ")
			}
		}
		return result
	}()
	daysLine := lipgloss.JoinHorizontal(lipgloss.Top, daysOfWeek...)

	firstDay := time.Date(cm.viewing.Year(), cm.viewing.Month(), 1, 0, 0, 0, 0, cm.viewing.Location())
	lastDay := firstDay.AddDate(0, 1, -1)

	// align calendar to start on Monday (make Sunday = 7)
	weekday := int(firstDay.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	dayLength := lipgloss.Width(styles.Base.Render("Mon"))
	emptySpace := strings.Repeat(" ", (dayLength+1)*(weekday-1))

	dates := ""
	datesLine := emptySpace
	for day := firstDay; !day.After(lastDay); day = day.AddDate(0, 0, 1) {
		isToday := day.Year() == cm.now.Year() && day.Month() == cm.now.Month() && day.Day() == cm.now.Day()
		isSelected := day.Year() == cm.selected.Year() && day.Month() == cm.selected.Month() && day.Day() == cm.selected.Day()

		var dayStr string
		if isSelected {
			dayStr = styles.SelectedDate.Width(dayLength).Render(fmt.Sprintf("%d", day.Day()))
		} else if isToday {
			dayStr = styles.Today.Width(dayLength).Render(fmt.Sprintf("%d", day.Day()))
		} else if utils.HasEvents(m.events, day) {
			dayStr = styles.DateWithEvent.Width(dayLength).Render(fmt.Sprintf("%d", day.Day()))
		} else {
			dayStr = styles.Base.Width(dayLength).Render(fmt.Sprintf("%d", day.Day()))
		}
		datesLine = lipgloss.JoinHorizontal(lipgloss.Top, datesLine, dayStr, " ")

		// break line at Sunday (weekday = 0)
		w := int(day.Weekday())
		if w == 0 {
			dates = lipgloss.JoinVertical(lipgloss.Top, dates, datesLine, "")
			datesLine = ""
		}
	}
	// add last line if not empty
	if datesLine != "" {
		dates = lipgloss.JoinVertical(lipgloss.Top, dates, datesLine)
	}

	// display events (if any) for the selected date
	dateKey := cm.selected.Format("2006-01-02")
	eventsBlock := ""

	if events, ok := m.events[dateKey]; ok && len(events) > 0 {
		headerText := "Events for " + cm.selected.Format("January 2, 2006")
		calWidth := lipgloss.Width(daysLine)
		eventsHeader := styles.EventHeader.Width(calWidth).Align(lipgloss.Center).Render(headerText)

		var rows []string

		for _, event := range events {
			timeLabel := "All Day"
			if event.Start != nil && event.Start.DateTime != "" {
				if t, err := time.Parse(time.RFC3339, event.Start.DateTime); err == nil {
					timeLabel = t.Format("15:04")
				}
			}
			timeStr := styles.TimeValue.Render(timeLabel)

			title := event.Summary
			if title == "" {
				title = "(No Title)"
			}

			minGap := 1
			available := calWidth - lipgloss.Width(timeStr) - minGap

			titleStr := utils.TruncateString(title, available, styles.Event)

			gap := max(available-lipgloss.Width(titleStr), minGap)

			row := lipgloss.JoinHorizontal(
				lipgloss.Top,
				timeStr,
				strings.Repeat(" ", gap),
				titleStr,
			)

			rows = append(rows, row)
		}

		eventsBlock = lipgloss.JoinVertical(
			lipgloss.Top,
			"",
			eventsHeader,
			"",
			lipgloss.JoinVertical(lipgloss.Top, rows...),
		)
	}

	helpText := lipgloss.PlaceHorizontal(m.screenWidth, lipgloss.Center, m.help.View(m.calendarViewKeys))

	calendarContent := lipgloss.PlaceHorizontal(m.screenWidth, lipgloss.Center, lipgloss.JoinVertical(
		lipgloss.Top,
		"",
		"",
		header,
		"",
		"",
		daysLine,
		dates,
		eventsBlock,
	))

	allContent := lipgloss.JoinVertical(
		lipgloss.Top,
		calendarContent,
		"",
		helpText,
	)

	return lipgloss.PlaceVertical(m.screenHeight, lipgloss.Top, allContent)
}
