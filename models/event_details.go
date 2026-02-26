package models

import (
	"math"
	"time"

	"github.com/NStefan002/tui-calendar/v2/styles"
	"github.com/NStefan002/tui-calendar/v2/utils"
	"google.golang.org/api/calendar/v3"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type eventDetailsModel struct {
	idx int // index of the currently selected event in the list

	viewport viewport.Model
}

func newDM() *eventDetailsModel {
	return &eventDetailsModel{
		idx: 0,

		viewport: viewport.New(0, 0),
	}
}

func (dm *eventDetailsModel) reset() {
	dm.idx = 0
	dm.viewport = viewport.New(0, 0)
}

func (dm *eventDetailsModel) view(m *model) string {
	dateKey := m.cm.selected.Format("2006-01-02")
	selectedEvents := m.events[dateKey]
	if len(selectedEvents) == 0 {
		return lipgloss.PlaceHorizontal(m.screenWidth, lipgloss.Center, "No events for this day.")
	}

	leftWidth, rightWidth := int(math.Round(float64(m.screenWidth)*0.4)), int(math.Round(float64(m.screenWidth)*0.6))

	selected := selectedEvents[dm.idx]

	titles := dm.getEventList(selectedEvents, leftWidth)
	left := lipgloss.JoinVertical(lipgloss.Left, titles...)

	dm.viewport.Width = rightWidth - 4 // account for padding
	title := dm.getTitle(selected, dm.viewport.Width)
	times := dm.getTimes(selected, dm.viewport.Width)
	location := dm.getLocation(selected, dm.viewport.Width)
	dm.viewport.Height = m.screenHeight - lipgloss.Height(title) - lipgloss.Height(times) - lipgloss.Height(location) - lipgloss.Height(m.help.View(m.eventDetailsViewKeys)) - 5
	dm.viewport.SetContent(dm.getDescription(selected, dm.viewport.Width))
	right := styles.Box.Render(lipgloss.JoinVertical(lipgloss.Left,
		title,
		dm.viewport.View(),
		"",
		times,
		location,
	))

	helpText := lipgloss.PlaceHorizontal(m.screenWidth, lipgloss.Center, m.help.View(m.eventDetailsViewKeys))

	top := lipgloss.PlaceHorizontal(m.screenWidth, lipgloss.Center, lipgloss.JoinHorizontal(lipgloss.Top, left, right))

	return lipgloss.PlaceVertical(m.screenHeight, lipgloss.Top, lipgloss.JoinVertical(lipgloss.Left, top, "", helpText))
}

func (dm *eventDetailsModel) getEventList(selectedEvents []*calendar.Event, maxWidth int) []string {
	titles := make([]string, len(selectedEvents))
	for i, event := range selectedEvents {
		title := event.Summary
		if title == "" {
			title = "[No Title]"
		}
		if i == dm.idx {
			titles[i] = styles.EventListSelected.Width(maxWidth).Render(utils.TruncateString(title, maxWidth-2, styles.EventListSelected))
		} else {
			titles[i] = styles.EventList.Width(maxWidth).Render(utils.TruncateString(title, maxWidth-2, styles.EventList))
		}
	}

	return titles
}

func (dm *eventDetailsModel) getTitle(event *calendar.Event, maxWidth int) string {
	title := event.Summary
	if title == "" {
		title = "[No Title]"
	}
	return lipgloss.PlaceHorizontal(maxWidth, lipgloss.Center, styles.DetailTitle.Width(maxWidth).Render(title))
}

func (dm *eventDetailsModel) getDescription(event *calendar.Event, maxWidth int) string {
	if event.Description == "" {
		return lipgloss.PlaceHorizontal(maxWidth, lipgloss.Center, styles.Description.Width(maxWidth).Render("[No description]"))
	}
	return lipgloss.PlaceHorizontal(maxWidth, lipgloss.Center, styles.Description.Width(maxWidth).Render(event.Description))
}

func (dm *eventDetailsModel) getTimes(event *calendar.Event, maxWidth int) string {
	if event.Start != nil && event.Start.Date != "" {
		return lipgloss.PlaceHorizontal(maxWidth, lipgloss.Center, styles.TimeValue.Render("All-day event"))
	}

	startStr, endStr := "", ""
	if event.Start != nil && event.Start.DateTime != "" {
		startTime, err := time.Parse(time.RFC3339, event.Start.DateTime)
		if err == nil {
			startStr = styles.TimeLabel.Render("Start: ") + styles.TimeValue.Render(startTime.Format("Mon, Jan 2 — 15:04"))
		}
	}
	if event.End != nil && event.End.DateTime != "" {
		endTime, err := time.Parse(time.RFC3339, event.End.DateTime)
		if err == nil {
			endStr = styles.TimeLabel.Render("End:   ") + styles.TimeValue.Render(endTime.Format("Mon, Jan 2 — 15:04"))
		}
	}
	return lipgloss.PlaceHorizontal(maxWidth, lipgloss.Center, lipgloss.JoinVertical(lipgloss.Left, startStr, endStr))
}

func (dm *eventDetailsModel) getLocation(event *calendar.Event, maxWidth int) string {
	if event.Location == "" {
		return ""
	}
	str := styles.LocationLabel.Render("Location:") + " " + styles.TimeValue.Render(event.Location)
	return lipgloss.PlaceHorizontal(maxWidth, lipgloss.Center, str)
}
