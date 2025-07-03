package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// styles for the calendar UI

var (
	baseStyle = lipgloss.NewStyle().
			Padding(0, 1).
			// Border(lipgloss.NormalBorder()).
			// BorderForeground(lipgloss.Color("240")).
			Background(lipgloss.Color("235"))

	headerStyle = baseStyle.Background(lipgloss.Color("39")).Bold(true).Align(lipgloss.Center)

	selectedDateStyle = baseStyle.Background(lipgloss.Color("45")).Bold(true)

	todayStyle = baseStyle.Background(lipgloss.Color("21")).Bold(true)

	eventHeaderStyle = baseStyle.Background(lipgloss.Color("33")).Bold(true)

	eventStyle = baseStyle.Foreground(lipgloss.Color("252"))
)
