package models

import (
	"github.com/charmbracelet/bubbles/key"
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
		{k.PrevDay, k.NextDay, k.PrevWeek, k.NextWeek},
		{k.PrevMonth, k.NextMonth},
		{k.ViewEvent, k.AddEvent},
		{k.Refresh},
		{k.Help, k.Quit},
	}
}

var calendarViewKeys = calendarViewKeyMap{
	PrevDay: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "previous day"),
	),
	NextDay: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "next day"),
	),
	PrevWeek: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "previous week"),
	),
	NextWeek: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "next week"),
	),
	PrevMonth: key.NewBinding(
		key.WithKeys("pageup", "pgup", "ctrl+u"),
		key.WithHelp("pgup/ctrl+u", "previous month"),
	),
	NextMonth: key.NewBinding(
		key.WithKeys("pagedown", "pgdown", "ctrl+d"),
		key.WithHelp("pgdown/ctrl+d", "next month"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
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
		key.WithKeys("enter"),
		key.WithHelp("enter", "view event details"),
	),
	AddEvent: key.NewBinding(
		key.WithKeys("a", "A"),
		key.WithHelp("a/A", "add new event"),
	),
}

type eventDetailsViewKeyMap struct {
	Quit       key.Binding
	Back       key.Binding
	Help       key.Binding
	ScrollDown key.Binding
	ScrollUp   key.Binding
	EditEvent  key.Binding
	AddEvent   key.Binding
}

func (k eventDetailsViewKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Back, k.Help}
}

func (k eventDetailsViewKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ScrollUp, k.ScrollDown},
		{k.EditEvent, k.AddEvent},
		{k.Back, k.Quit, k.Help},
	}
}

var eventDetailsViewKeys = eventDetailsViewKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back to previous view"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	ScrollDown: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "scroll down"),
	),
	ScrollUp: key.NewBinding(
		key.WithKeys("up", "k"),
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
}

type editEventViewKeyMap struct {
	Quit key.Binding
	Back key.Binding
	Help key.Binding
}

func (k editEventViewKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Back, k.Help}
}

func (k editEventViewKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Back, k.Quit, k.Help},
	}
}

var editEventViewKeys = editEventViewKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back to previous view"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
}

type addEventViewKeyMap struct {
	Quit     key.Binding
	Back     key.Binding
	Help     key.Binding
	Next     key.Binding
	Previous key.Binding
	Submit   key.Binding
}

func (k addEventViewKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Back, k.Help}
}

func (k addEventViewKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Next, k.Previous},
		{k.Submit},
		{k.Back, k.Quit, k.Help},
	}
}

var addEventViewKeys = addEventViewKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back to previous view"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Next: key.NewBinding(
		key.WithKeys("tab", "down", "ctrl+n"),
		key.WithHelp("tab/↓/ctrl+n", "next field"),
	),
	Previous: key.NewBinding(
		key.WithKeys("shift+tab", "up", "ctrl+p"),
		key.WithHelp("shift+tab/↑/ctrl+p", "previous field"),
	),
	Submit: key.NewBinding(
		key.WithKeys("ctrl+s", "enter"),
		key.WithHelp("ctrl+s/enter", "submit"),
	),
}
