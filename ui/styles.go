package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// styles for the calendar UI

var (
	baseStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#DADADA")).
			Background(lipgloss.Color("#1E1E1E")).
			Padding(0, 1)

	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#3A3A3A")).
			Padding(1, 2).
			Bold(true).
			Align(lipgloss.Center)

	selectedDateStyle = baseStyle.
				Foreground(lipgloss.Color("#1E1E1E")).
				Background(lipgloss.Color("#FFD700")).
				Bold(true)

	todayStyle = baseStyle.
			Background(lipgloss.Color("#005F87")).
			Bold(true)

	dateWithEventStyle = baseStyle.
				Foreground(lipgloss.Color("#F0F0F0")).
				Background(lipgloss.Color("#444444")).
				Bold(true)

	eventHeaderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#005F5F")).
				Bold(true).
				Padding(0, 1)

	eventStyle = baseStyle.
			Foreground(lipgloss.Color("#C0C0C0"))

	eventDetailsStyle = lipgloss.NewStyle().
				Padding(1, 2).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#5F5F5F")).
				Width(50)

	eventListStyle = baseStyle.
			Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#3A3A3A"))

	eventListSelectedStyle = eventListStyle.
				Background(lipgloss.Color("#5F00AF")).
				Foreground(lipgloss.Color("#FFFFFF")).
				Bold(true)

	activeTextinputStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#005F87")).
				Padding(0, 1).
				Bold(true)

	inactiveTextinputStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#888888")).
				Background(lipgloss.Color("#1E1E1E")).
				Padding(0, 1)
)
