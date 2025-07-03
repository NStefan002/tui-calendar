package main

import (
	"log"
	"tui-calendar/google"
	"tui-calendar/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	srv, err := google.GetClient()
	if err != nil {
		log.Fatalf("Failed to get calendar client: %v", err)
	}

	p := tea.NewProgram(ui.CreateModel(srv))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
