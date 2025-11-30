package models

import (
	"fmt"
	"tui-calendar/google"
	"tui-calendar/utils"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"google.golang.org/api/calendar/v3"
)

type errMsg struct{ error }

type eventsMsg map[string][]*calendar.Event

func (m model) Init() tea.Cmd {
	m.loading = true
	return tea.Batch(
		m.spinner.Tick, // start spinner on app load
		func() tea.Msg {
			events, err := google.FetchEvents(m.calendarService, m.cm.viewing)
			if err != nil {
				return errMsg{err}
			}
			return eventsMsg(events)
		},
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.screenWidth = msg.Width
		m.screenHeight = msg.Height

	// custom message for events
	case eventsMsg:
		m.events = msg
		m.loading = false

	// custom message for events
	case errMsg:
		m.loading = false
		m.errMessage = msg.Error()

	case tea.KeyMsg:
		switch m.viewMode {
		case calendarView:
			switch {
			case key.Matches(msg, m.calendarViewKeys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.calendarViewKeys.PrevDay):
				m.cm.selected = m.cm.selected.AddDate(0, 0, -1) // go to previous day
			case key.Matches(msg, m.calendarViewKeys.NextDay):
				m.cm.selected = m.cm.selected.AddDate(0, 0, 1) // go to next day
			case key.Matches(msg, m.calendarViewKeys.PrevWeek):
				m.cm.selected = m.cm.selected.AddDate(0, 0, -7) // go to previous week
			case key.Matches(msg, m.calendarViewKeys.NextWeek):
				m.cm.selected = m.cm.selected.AddDate(0, 0, 7) // go to next week
			case key.Matches(msg, m.calendarViewKeys.PrevMonth):
				m.cm.selected = m.cm.selected.AddDate(0, -1, 0) // go to previous month
			case key.Matches(msg, m.calendarViewKeys.NextMonth):
				m.cm.selected = m.cm.selected.AddDate(0, 1, 0) // go to next month
			case key.Matches(msg, m.calendarViewKeys.Refresh):
				m.loading = true
				return m, tea.Batch(
					m.spinner.Tick, // start spinner
					func() tea.Msg {
						events, err := google.FetchEvents(m.calendarService, m.cm.viewing)
						if err != nil {
							return errMsg{err}
						}
						return eventsMsg(events)
					},
				)
			case key.Matches(msg, m.calendarViewKeys.Help):
				m.help.ShowAll = !m.help.ShowAll
			case key.Matches(msg, m.calendarViewKeys.ViewEvent):
				m.viewMode = eventDetailsView
				m.lastViewMode = calendarView
				m.help.ShowAll = false
			case key.Matches(msg, m.calendarViewKeys.AddEvent):
				m.viewMode = addEventView
				m.lastViewMode = calendarView
				m.help.ShowAll = false
				m.am.selectedDate = m.cm.selected
				m.am.resetForm()
			}

		case eventDetailsView:
			switch {
			case key.Matches(msg, m.eventDetailsViewKeys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.eventDetailsViewKeys.Back):
				m.viewMode = calendarView
			case key.Matches(msg, m.eventDetailsViewKeys.Help):
				m.help.ShowAll = !m.help.ShowAll
			case key.Matches(msg, m.eventDetailsViewKeys.ScrollDown):
				if len(m.events) > 0 && m.dm.idx < len(m.events[m.cm.selected.Format("2006-01-02")])-1 {
					m.dm.idx++
				}
			case key.Matches(msg, m.eventDetailsViewKeys.ScrollUp):
				if m.dm.idx > 0 {
					m.dm.idx--
				}
			case key.Matches(msg, m.eventDetailsViewKeys.EditEvent):
				m.viewMode = editEventView
				m.lastViewMode = eventDetailsView
				m.help.ShowAll = false
			case key.Matches(msg, m.eventDetailsViewKeys.AddEvent):
				m.viewMode = addEventView
				m.lastViewMode = eventDetailsView
				m.help.ShowAll = false
			case key.Matches(msg, m.eventDetailsViewKeys.DeleteEvent):
				event := m.events[m.cm.selected.Format("2006-01-02")][m.dm.idx]
				err := google.DeleteEvent(m.calendarService, event.Id)
				if err != nil {
					m.errMessage = fmt.Sprintf("Failed to delete event: %v", err)
					m.viewMode = m.lastViewMode
					return m, nil
				}
				// return to calendar view after deleting event
				m.viewMode = calendarView
				m.help.ShowAll = false
				m.dm.idx = 0 // reset event index
				// refresh events after deleting event
				m.loading = true
				return m, tea.Batch(
					m.spinner.Tick, // start spinner
					func() tea.Msg {
						events, err := google.FetchEvents(m.calendarService, m.cm.viewing)
						if err != nil {
							return errMsg{err}
						}
						return eventsMsg(events)
					},
				)
			}

		case editEventView:
			switch {
			case key.Matches(msg, m.editEventViewKeys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.editEventViewKeys.Back):
				m.viewMode = m.lastViewMode
				m.help.ShowAll = false
			case key.Matches(msg, m.editEventViewKeys.Help):
				m.help.ShowAll = !m.help.ShowAll
			}

		case addEventView:
			switch {
			case key.Matches(msg, m.addEventViewKeys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.addEventViewKeys.Back):
				m.viewMode = m.lastViewMode
				m.help.ShowAll = false
			case key.Matches(msg, m.addEventViewKeys.Help):
				m.help.ShowAll = !m.help.ShowAll
			case key.Matches(msg, m.addEventViewKeys.Check) && m.am.checkBoxFocused():
				m.am.toggleAllDay()
			case key.Matches(msg, m.addEventViewKeys.MinuteUp) && m.am.timeFieldFocused():
				m.am.changeMinutes(+1)
			case key.Matches(msg, m.addEventViewKeys.MinuteDown) && m.am.timeFieldFocused():
				m.am.changeMinutes(-1)
			case key.Matches(msg, m.addEventViewKeys.HourUp) && m.am.timeFieldFocused():
				m.am.changeHours(+1)
			case key.Matches(msg, m.addEventViewKeys.HourDown) && m.am.timeFieldFocused():
				m.am.changeHours(-1)
			case key.Matches(msg, m.addEventViewKeys.Next):
				m.am.changeFocus(+1)
			case key.Matches(msg, m.addEventViewKeys.Previous):
				m.am.changeFocus(-1)
			case key.Matches(msg, m.addEventViewKeys.Submit):
				event, err := m.am.submit()
				if err != nil {
					m.errMessage = fmt.Sprintf("Failed to submit form: %v", err)
					m.viewMode = m.lastViewMode
					return m, nil
				}
				_, err = google.CreateEvent(m.calendarService, event)
				if err != nil {
					m.errMessage = fmt.Sprintf("Failed to create event: %v", err)
					m.viewMode = m.lastViewMode
					return m, nil
				}
				// return to calendar view after creating event
				m.viewMode = calendarView
				m.am.resetForm()
				m.help.ShowAll = false
				// refresh events after creating new event
				m.loading = true
				return m, tea.Batch(
					m.spinner.Tick, // start spinner
					func() tea.Msg {
						events, err := google.FetchEvents(m.calendarService, m.cm.viewing)
						if err != nil {
							return errMsg{err}
						}
						return eventsMsg(events)
					},
				)
			default:
				cmds := make([]tea.Cmd, 3)
				m.am.titleInput, cmds[0] = m.am.titleInput.Update(msg)
				m.am.descriptionInput, cmds[1] = m.am.descriptionInput.Update(msg)
				m.am.locationInput, cmds[2] = m.am.locationInput.Update(msg)

				return m, tea.Batch(cmds...)
			}
		}

	// spinner update
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	if m.cm.selected.Month() != m.cm.viewing.Month() || m.cm.selected.Year() != m.cm.viewing.Year() {
		m.cm.viewing = m.cm.selected // update viewing month if selected date is not in current viewing month
	}

	return m, nil
}

func (m model) View() string {
	if m.loading {
		return utils.CenterText(m.spinner.View()+" Loading calendar events...", m.screenWidth)
	}

	if len(m.errMessage) > 0 {
		return "There was an error: " + m.errMessage
	}

	switch m.viewMode {
	case calendarView:
		return m.cm.view(&m)
	case eventDetailsView:
		return m.dm.view(&m)
	case editEventView:
		return m.em.view()
	case addEventView:
		return m.am.view(&m)
	default:
		return "ERROR"
	}
}
