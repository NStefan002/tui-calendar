package models

import (
	"strings"
	"time"
	"tui-calendar/styles"
	"tui-calendar/utils"

	"github.com/charmbracelet/lipgloss"
	"google.golang.org/api/calendar/v3"
)

type eventDetailsModel struct {
	idx int // index of the currently selected event in the list
}

func newDM() *eventDetailsModel {
	return &eventDetailsModel{
		idx: 0,
	}
}

func (dm *eventDetailsModel) view(selectedDate time.Time, events map[string][]*calendar.Event, scrWidth, scrHeight int) string {
	dateKey := selectedDate.Format("2006-01-02")
	selectedEvents := events[dateKey]
	if len(selectedEvents) == 0 {
		return utils.CenterText("No events for this day.", scrWidth)
	}

	selected := selectedEvents[dm.idx]

	// left column: event titles
	var list strings.Builder
	for i, event := range selectedEvents {
		title := event.Summary
		if title == "" {
			title = "[No Title]"
		}
		if i == dm.idx {
			list.WriteString(styles.EventListSelected.Render(title) + "\n")
		} else {
			list.WriteString(styles.EventList.Render(title) + "\n")
		}
	}
	leftCol := styles.Box.Width(30).Render(list.String())

	// right column: selected event details
	var right strings.Builder

	// title
	eventTitle := selected.Summary
	if eventTitle == "" {
		eventTitle = "[No Title]"
	}
	right.WriteString(utils.CenterText(styles.DetailTitle.Render(eventTitle), 50) + "\n\n")

	// times
	var startStr, endStr string
	if selected.Start != nil && selected.Start.DateTime != "" {
		startTime, err := time.Parse(time.RFC3339, selected.Start.DateTime)
		if err == nil {
			startStr = styles.TimeLabel.Render("Start: ") + styles.TimeValue.Render(startTime.Format("Mon, Jan 2 — 15:04"))
		}
	}
	if selected.End != nil && selected.End.DateTime != "" {
		endTime, err := time.Parse(time.RFC3339, selected.End.DateTime)
		if err == nil {
			endStr = styles.TimeLabel.Render("End:   ") + styles.TimeValue.Render(endTime.Format("Mon, Jan 2 — 15:04"))
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
	right.WriteString("\n" + styles.Description.Render(desc))

	rightCol := styles.Box.Width(50).Render(right.String())

	// side-by-side layout
	main := lipgloss.JoinHorizontal(lipgloss.Top, leftCol, rightCol)

	// footer
	footer := styles.FormFooter.Render("[j/k] Navigate  [esc] Back")

	return utils.CenterText(main, scrWidth) + "\n\n" + utils.CenterText(footer, scrWidth)
}
