package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// styles for the calendar UI

var (
	Base = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DADADA")).
		Background(lipgloss.Color("#1E1E1E")).
		Padding(0, 1)

	Header = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#3A3A3A")).
		Padding(1, 2).
		Bold(true).
		Align(lipgloss.Center)

	SelectedDate = Base.
			Foreground(lipgloss.Color("#1E1E1E")).
			Background(lipgloss.Color("#FFD700")).
			Bold(true)

	Today = Base.
		Background(lipgloss.Color("#005F87")).
		Bold(true)

	DateWithEvent = Base.
			Foreground(lipgloss.Color("#F0F0F0")).
			Background(lipgloss.Color("#444444")).
			Bold(true)

	EventHeader = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#005F5F")).
			Bold(true).
			Padding(0, 1)

	Event = Base.
		Foreground(lipgloss.Color("#C0C0C0"))

	EventDetails = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#5F5F5F")).
			Width(50)

	EventList = Base.
			Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#3A3A3A"))

	EventListSelected = EventList.
				Background(lipgloss.Color("#5F00AF")).
				Foreground(lipgloss.Color("#FFFFFF")).
				Bold(true)

	ActiveTextinput = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#005F87")).
			Padding(0, 1).
			Bold(true)

	InactiveTextinput = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#888888")).
				Background(lipgloss.Color("#1E1E1E")).
				Padding(0, 1)

	Box = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Margin(1, 0).
		BorderForeground(lipgloss.Color("240"))

	FieldLabel = lipgloss.NewStyle().
			Foreground(lipgloss.Color("250")).
			PaddingRight(2).
			Width(15)

	FormFooter = lipgloss.NewStyle().
			Background(lipgloss.Color("236")).
			Foreground(lipgloss.Color("250")).
			Padding(0, 1).
			Align(lipgloss.Center)

	DetailTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#005F87")).
			Padding(0, 1)

	TimeLabel = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")).
			Bold(true)

	TimeValue = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			PaddingLeft(1)

	Description = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#C0C0C0")).
			PaddingTop(1)

	LocationLabel = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")).
			Bold(true)

	LocationValue = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))
)
