package models

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/bubbles/help"
)

type calendarViewKeyMap struct {
	PrevDay   key.Binding
	NextDay   key.Binding
	PrevWeek  key.Binding
	NextWeek  key.Binding
	PrevMonth key.Binding
	NextMonth key.Binding
	Quit      key.Binding
	Help      key.Binding
	Refresh   key.Binding
	ViewEvent key.Binding
	AddEvent  key.Binding
}

func (k calendarViewKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help}
}

func (k calendarViewKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PrevDay, k.NextDay, k.PrevWeek, k.NextWeek, k.PrevMonth, k.NextMonth},
		{k.ViewEvent, k.AddEvent},
		{k.Refresh, k.Help, k.Quit},
	}
}

var calendarViewKeys = calendarViewKeyMap{
	PrevDay: key.NewBinding(
		key.WithKeys(tea.KeyLeft.String(), "h"),
		key.WithHelp("←/h", "previous day"),
	),
	NextDay: key.NewBinding(
		key.WithKeys(tea.KeyRight.String(), "l"),
		key.WithHelp("→/l", "next day"),
	),
	PrevWeek: key.NewBinding(
		key.WithKeys(tea.KeyUp.String(), "k"),
		key.WithHelp("↑/k", "previous week"),
	),
	NextWeek: key.NewBinding(
		key.WithKeys(tea.KeyDown.String(), "j"),
		key.WithHelp("↓/j", "next week"),
	),
	PrevMonth: key.NewBinding(
		key.WithKeys(tea.KeyPgUp.String(), tea.KeyCtrlU.String()),
		key.WithHelp("pgup/ctrl+u", "previous month"),
	),
	NextMonth: key.NewBinding(
		key.WithKeys(tea.KeyPgDown.String(), tea.KeyCtrlD.String()),
		key.WithHelp("pgdown/ctrl+d", "next month"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", tea.KeyCtrlC.String()),
		key.WithHelp("q/ctrl+c", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Refresh: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refresh events"),
	),
	ViewEvent: key.NewBinding(
		key.WithKeys(tea.KeyEnter.String()),
		key.WithHelp("enter", "view event details"),
	),
	AddEvent: key.NewBinding(
		key.WithKeys("a", "A"),
		key.WithHelp("a/A", "add new event"),
	),
}

type eventDetailsViewKeyMap struct {
	Quit        key.Binding
	Back        key.Binding
	Help        key.Binding
	ScrollDown  key.Binding
	ScrollUp    key.Binding
	EditEvent   key.Binding
	AddEvent    key.Binding
	DeleteEvent key.Binding
}

func (k eventDetailsViewKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Back, k.Help}
}

func (k eventDetailsViewKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ScrollUp, k.ScrollDown},
		{k.EditEvent, k.AddEvent, k.DeleteEvent},
		{k.Back, k.Quit, k.Help},
	}
}

var eventDetailsViewKeys = eventDetailsViewKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", tea.KeyCtrlC.String()),
		key.WithHelp("q/ctrl+c", "quit"),
	),
	Back: key.NewBinding(
		key.WithKeys(tea.KeyEsc.String()),
		key.WithHelp("esc", "back to previous view"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	ScrollDown: key.NewBinding(
		key.WithKeys(tea.KeyDown.String(), "j"),
		key.WithHelp("↓/j", "scroll down"),
	),
	ScrollUp: key.NewBinding(
		key.WithKeys(tea.KeyUp.String(), "k"),
		key.WithHelp("↑/k", "scroll up"),
	),
	EditEvent: key.NewBinding(
		key.WithKeys("e", "E"),
		key.WithHelp("e/E", "edit event"),
	),
	AddEvent: key.NewBinding(
		key.WithKeys("a", "A"),
		key.WithHelp("a/A", "add new event"),
	),
	DeleteEvent: key.NewBinding(
		key.WithKeys("D", tea.KeyDelete.String()),
		key.WithHelp("D/del", "delete event"),
	),
}

type addEventViewKeyMap struct {
	Quit       key.Binding
	Back       key.Binding
	Help       key.Binding
	MinuteUp   key.Binding
	MinuteDown key.Binding
	HourUp     key.Binding
	HourDown   key.Binding
	Check      key.Binding
	Next       key.Binding
	Previous   key.Binding
	Submit     key.Binding
}

func (k addEventViewKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Back, k.Help}
}

func (k addEventViewKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Next, k.Previous},
		{k.MinuteUp, k.MinuteDown, k.HourUp, k.HourDown},
		{k.Check, k.Submit},
		{k.Back, k.Quit, k.Help},
	}
}

var addEventViewKeys = addEventViewKeyMap{
	Quit: key.NewBinding(
		key.WithKeys(tea.KeyCtrlC.String()),
		key.WithHelp("ctrl+c", "quit"),
	),
	Back: key.NewBinding(
		key.WithKeys(tea.KeyEsc.String()),
		key.WithHelp("esc", "back to previous view"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	MinuteUp: key.NewBinding(
		key.WithKeys(tea.KeyCtrlUp.String(), "K"),
		key.WithHelp("ctrl+↑/K", "increase minutes"),
	),
	MinuteDown: key.NewBinding(
		key.WithKeys(tea.KeyCtrlDown.String(), "J"),
		key.WithHelp("ctrl+↓/J", "decrease minutes"),
	),
	HourUp: key.NewBinding(
		key.WithKeys("alt+up", "k"),
		key.WithHelp("alt+↑/k", "increase hours"),
	),
	HourDown: key.NewBinding(
		key.WithKeys("alt+down", "j"),
		key.WithHelp("alt+↓/j", "decrease hours"),
	),
	Check: key.NewBinding(
		key.WithKeys(tea.KeyEnter.String(), tea.KeySpace.String()),
		key.WithHelp("enter/space", "toggle all-day event"),
	),
	Next: key.NewBinding(
		key.WithKeys(tea.KeyTab.String(), tea.KeyDown.String(), tea.KeyCtrlN.String()),
		key.WithHelp("tab/↓/ctrl+n", "next field"),
	),
	Previous: key.NewBinding(
		key.WithKeys(tea.KeyShiftTab.String(), tea.KeyUp.String(), tea.KeyCtrlP.String()),
		key.WithHelp("shift+tab/↑/ctrl+p", "previous field"),
	),
	Submit: key.NewBinding(
		key.WithKeys(tea.KeyCtrlS.String()),
		key.WithHelp("ctrl+s", "submit"),
	),
}
