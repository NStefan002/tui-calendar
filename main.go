package main

import (
	"log"
	"tui-calendar/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.CreateModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
